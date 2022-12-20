package service

import (
	"context"

	health "google.golang.org/grpc/health/grpc_health_v1"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

type Database interface {
	Ping(ctx context.Context) error
	GetUser(ctx context.Context, id string) (*userv1.User, error)
	StoreUser(ctx context.Context, user *userv1.User) error
}

type Service struct {
	DB Database

	userv1.UnimplementedUserServiceServer
	health.UnimplementedHealthServer
}

func NewService(db Database) *Service {
	return &Service{DB: db}
}

func (s *Service) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	res := &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}
	return res, nil
}
