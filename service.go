package main

import (
	"context"
	"time"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

type Service struct {
	db Database

	userv1.UnimplementedUserServiceServer
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
	user := req.GetUser()

	user.Created = time.Now().Unix()

	if err := s.db.StoreUser(ctx, user); err != nil {
		return nil, err
	}

	res := &userv1.CreateUserResponse{
		User: user,
	}

	return res, nil
}
