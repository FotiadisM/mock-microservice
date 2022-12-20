package service

import (
	"context"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

func (s *Service) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := s.DB.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	res := &userv1.GetUserResponse{
		User: user,
	}

	return res, nil
}
