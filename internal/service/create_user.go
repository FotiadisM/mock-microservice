package service

import (
	"context"
	"time"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
	"github.com/google/uuid"
)

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

	if err := s.DB.StoreUser(ctx, user); err != nil {
		return nil, err
	}

	res := &userv1.CreateUserResponse{
		User: user,
	}

	return res, nil
}
