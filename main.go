package main

import (
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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

var (
	addr        = flag.String("addr", ":8080", "grpc listening address")
	metricsAddr = flag.String("metrics_addr", ":9090", "metrics listening address")
	production  = flag.Bool("prod", false, "production mode")
)

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.Fatal("net.Listen failed", zap.Error(err))
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_opentracing.UnaryServerInterceptor(),
		),
	)

	grpc_prometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())

	if !*production {
		reflection.Register(server)
	}

	db := newInMemoryDB()
	svc := NewService(db)
	userv1.RegisterUserServiceServer(server, svc)

	go func() {
		logger.Info("server is listening", zap.String("port", *addr))
		if err := server.Serve(lis); err != nil {
			logger.Fatal("grpc.Serve failed", zap.Error(err))
		}
	}()

	go func() {
		logger.Info("metrics available", zap.String("port", *addr))
		if err := http.ListenAndServe(*metricsAddr, http.DefaultServeMux); err != nil {
			logger.Fatal("http.ListenAndServe failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	server.GracefulStop()
	logger.Info("server exited properly")
}
