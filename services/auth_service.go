package services

import (
	"context"

	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/response"
)

type AuthService interface {
	Login(ctx context.Context, param params.Auth) (response.LoginResponse, error)
}
