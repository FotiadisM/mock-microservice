// Package server provides utilities for scaffolding an opinionated
// grpc server with a http gateway.

package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcopenmetrics "github.com/grpc-ecosystem/go-grpc-middleware/providers/openmetrics/v2"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/providers/opentracing/v2"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/providers/zap/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/tracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/findit-it/users-svc/pkg/grpc/health"
)

type Server interface {
	Start() error
	RegisterService(regFunc func(s *grpc.Server, m *runtime.ServeMux))
	AwaitTermination(ctx context.Context)
}

type Config struct {
	GrpcAddr     string
	HttpAddr     string
	Reflection   bool
	Logger       *zap.Logger
	HealthServer grpc_health_v1.HealthServer
	Options      []grpc.ServerOption

	LoggingOptions  []logging.Option
	TracingOptions  []grpcopentracing.Option
	MetricsOptions  []grpcopenmetrics.ServerMetricsOption
	Recoveryoptions []recovery.Option
	MuxOptions      []runtime.ServeMuxOption

	StreamInterceptors []grpc.StreamServerInterceptor
	UnaryInterceptors  []grpc.UnaryServerInterceptor
}

func NewDefaultConfig(logger *zap.Logger, hs grpc_health_v1.HealthServer) *Config {
	return &Config{
		GrpcAddr:           ":8080",
		HttpAddr:           ":9090",
		Reflection:         false,
		Logger:             logger,
		HealthServer:       hs,
		Options:            []grpc.ServerOption{},
		LoggingOptions:     []logging.Option{},
		TracingOptions:     []grpcopentracing.Option{},
		MetricsOptions:     []grpcopenmetrics.ServerMetricsOption{},
		Recoveryoptions:    []recovery.Option{},
		MuxOptions:         []runtime.ServeMuxOption{},
		StreamInterceptors: []grpc.StreamServerInterceptor{},
		UnaryInterceptors:  []grpc.UnaryServerInterceptor{},
	}
}

func (c *Config) Build() (Server, error) {
	serverMetrics := grpcopenmetrics.NewServerMetrics(c.MetricsOptions...)

	ui := []grpc.UnaryServerInterceptor{
		tracing.UnaryServerInterceptor(grpcopentracing.InterceptorTracer(c.TracingOptions...)),
		grpcopenmetrics.UnaryServerInterceptor(serverMetrics),
		logging.UnaryServerInterceptor(grpczap.InterceptorLogger(c.Logger), c.LoggingOptions...),
		recovery.UnaryServerInterceptor(c.Recoveryoptions...),
		validator.UnaryServerInterceptor(false),
	}
	ui = append(ui, c.UnaryInterceptors...)
	si := []grpc.StreamServerInterceptor{
		tracing.StreamServerInterceptor(grpcopentracing.InterceptorTracer(c.TracingOptions...)),
		grpcopenmetrics.StreamServerInterceptor(serverMetrics),
		logging.StreamServerInterceptor(grpczap.InterceptorLogger(c.Logger), c.LoggingOptions...),
		recovery.StreamServerInterceptor(c.Recoveryoptions...),
		validator.StreamServerInterceptor(false),
	}
	si = append(si, c.StreamInterceptors...)
	opts := append(c.Options,
		grpc.ChainUnaryInterceptor(ui...),
		grpc.ChainStreamInterceptor(si...),
	)

	grpcserver := grpc.NewServer(opts...)

	if c.Reflection {
		reflection.Register(grpcserver)
	}
	grpc_health_v1.RegisterHealthServer(grpcserver, c.HealthServer)

	c.MuxOptions = append(c.MuxOptions,
		runtime.WithHealthzEndpoint(health.NewHealthClient(c.HealthServer)),
	)
	mux := runtime.NewServeMux(c.MuxOptions...)
	if err := mux.HandlePath(http.MethodGet, "/metrics", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	}); err != nil {
		return nil, err
	}
	httpserver := &http.Server{Addr: c.HttpAddr, Handler: mux}

	gs := &server{
		config:        c,
		grpcserver:    grpcserver,
		listener:      nil,
		httpserver:    httpserver,
		serverMetrics: serverMetrics,
		mux:           mux,
	}

	return gs, nil
}

type server struct {
	config        *Config
	grpcserver    *grpc.Server
	listener      net.Listener
	httpserver    *http.Server
	serverMetrics *grpcopenmetrics.ServerMetrics
	mux           *runtime.ServeMux
}

func (s *server) Start() error {
	// Initialize metrics after all services have been registered
	s.serverMetrics.InitializeMetrics(s.grpcserver)

	var err error
	s.listener, err = net.Listen("tcp", s.config.GrpcAddr)
	if err != nil {
		return err
	}

	s.config.Logger.Info("grpc server is listening", zap.String("port", s.config.GrpcAddr))
	go func() {
		if err := s.grpcserver.Serve(s.listener); err != nil {
			s.config.Logger.Fatal("grpc serve failed", zap.Error(err))
		}
	}()

	s.config.Logger.Info("http server is listening", zap.String("port", s.config.HttpAddr))
	go func() {
		if err := s.httpserver.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				s.config.Logger.Fatal("http serve failed", zap.Error(err))
			}
		}
	}()

	return nil
}

func (s *server) RegisterService(regFunc func(s *grpc.Server, m *runtime.ServeMux)) {
	regFunc(s.grpcserver, s.mux)
}

func (s *server) AwaitTermination(ctx context.Context) {
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM)
	<-interruptSignal

	s.httpserver.Shutdown(ctx) //nolint:errcheck
	s.grpcserver.GracefulStop()
	s.listener.Close() //nolint:errcheck
}
