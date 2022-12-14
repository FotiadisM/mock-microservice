package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	health "google.golang.org/grpc/health/grpc_health_v1"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

type Database interface {
	Ping(ctx context.Context) error
	GetUser(ctx context.Context, id string) (*userv1.User, error)
	StoreUser(ctx context.Context, user *userv1.User) error
}

type Service struct {
	db Database

	userv1.UnimplementedUserServiceServer
	health.UnimplementedHealthServer
}

func NewService(db Database) userv1.UserServiceServer {
	return &Service{db: db}
}

func (s *Service) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := s.db.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	res := &userv1.GetUserResponse{
		User: user,
	}

	return res, nil
}

func (s *Service) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	user := &userv1.User{
		Id:          uuid.String(),
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		Email:       req.GetEmail(),
		Created:     time.Now().Unix(),
		PhoneNumber: req.GetPhoneNumber(),
	}

	if err := s.db.StoreUser(ctx, user); err != nil {
		return nil, err
	}

	res := &userv1.CreateUserResponse{
		User: user,
	}

	return res, nil
}

func (s *Service) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	res := &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}
	return res, nil
}
