package services

import (
	"context"

	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/response"
)

type UserService interface {
	Register(ctx context.Context, param params.User) (response.UserResponse, error)
}
