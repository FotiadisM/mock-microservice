package main

import (
	"context"
	"time"

	"google.golang.org/genproto/googleapis/type/datetime"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

type service struct {
	db Database

	userv1.UnimplementedUserServiceServer
}

func newService(db Database) *service {
	return &service{db: db}
}

func (s *service) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := s.db.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	res := &userv1.GetUserResponse{
		User: user,
	}

	return res, nil
}

func (s *service) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	user := req.GetUser()

	now := time.Now()
	user.Created = &datetime.DateTime{
		Year:       int32(now.Year()),
		Month:      int32(now.Month()),
		Day:        int32(now.Day()),
		Hours:      int32(now.Hour()),
		Minutes:    int32(now.Minute()),
		Seconds:    int32(now.Second()),
		Nanos:      0,
		TimeOffset: nil,
	}

	if err := s.db.StoreUser(ctx, user); err != nil {
		return nil, err
	}

	res := &userv1.CreateUserResponse{
		User: user,
	}

	return res, nil
}
