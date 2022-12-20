package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
	"github.com/findit-it/users-svc/internal/service"
	"github.com/findit-it/users-svc/pkg/grpc/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	grpc_addr = flag.String("grpc_addr", ":8080", "grpc listening address")
	http_addr = flag.String("http_addr", ":9090", "http listening address")
)

func main() {
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db := newInMemoryDB()
	svc := service.NewService(db)

	serverConfig := server.NewDefaultConfig(logger, svc)
	serverConfig.GrpcAddr = *grpc_addr
	serverConfig.HttpAddr = *http_addr
	serverConfig.Reflection = true
	server, err := serverConfig.Build()
	if err != nil {
		logger.Fatal("failed to build server", zap.Error(err))
	}

	ctx := context.Background()
	server.RegisterService(func(s *grpc.Server, m *runtime.ServeMux) {
		userv1.RegisterUserServiceServer(s, svc)
		if err := userv1.RegisterUserServiceHandlerServer(ctx, m, svc); err != nil {
			logger.Fatal("failed to register server", zap.Error(err))
		}
	})

	if err := server.Start(); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}

	server.AwaitTermination(ctx)
	logger.Info("server exited properly")
}
