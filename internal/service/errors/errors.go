package errors

import (
	"github.com/FotiadisM/mock-microservice/pkg/grpc/errors"
	"google.golang.org/grpc/codes"
)

var ErrEmailExists = errors.NewDetailsError(codes.AlreadyExists, "EMAIL_NOT_UNIQUE", "the email provided is already in use")
