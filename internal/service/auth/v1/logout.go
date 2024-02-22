package authv1

import (
	"context"

	authv1 "github.com/FotiadisM/mock-microservice/api/gen/go/auth/v1"
	"github.com/FotiadisM/mock-microservice/pkg/ilog"
)

func (s *Service) Logout(ctx context.Context, _ *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	log := ilog.FromContext(ctx)
	log.Info("HELLO")

	panic("oh no this is terrible")
}
