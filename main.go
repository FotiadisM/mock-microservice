package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

var (
	grpc_addr  = flag.String("grpc_addr", ":8080", "grpc listening address")
	http_addr  = flag.String("http_addr", ":9090", "http listening address")
	production = flag.Bool("prod", false, "production mode")
)

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db := newInMemoryDB()
	svc := NewService(db)

	lis, err := net.Listen("tcp", *grpc_addr)
	if err != nil {
		logger.Fatal("net.Listen failed", zap.Error(err))
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
		),
	)
	if !*production {
		reflection.Register(grpcServer)
	}
	userv1.RegisterUserServiceServer(grpcServer, svc)

	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	grpc_health_v1.RegisterHealthServer(grpcServer, svc.(*Service))

	go func() {
		logger.Info("grpc server is listening", zap.String("port", *grpc_addr))
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal("grpc.Serve failed", zap.Error(err))
		}
	}()

	hc := &healthClient{svc: svc.(*Service)}
	mux := runtime.NewServeMux(runtime.WithHealthEndpointAt(hc, "/healthz"))
	if err = mux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	}); err != nil {
		logger.Fatal("error registering metrics handler", zap.Error(err))
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = userv1.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, *grpc_addr, opts); err != nil {
		logger.Fatal("error registering http gateway", zap.Error(err))
	}

	httpServer := http.Server{Addr: *http_addr, Handler: mux}
	go func() {
		logger.Info("http server is listening", zap.String("port", *http_addr))
		if err = httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.Fatal("http.ListenAndServe failed", zap.Error(err))
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	grpcServer.GracefulStop()
	httpServer.Shutdown(context.Background()) //nolint:errcheck
	logger.Info("server exited properly")
}

// healthClient is an in-process HealthClient
type healthClient struct {
	svc grpc_health_v1.HealthServer

	grpc_health_v1.HealthClient
}

func (c *healthClient) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	return c.svc.Check(ctx, req)
}
