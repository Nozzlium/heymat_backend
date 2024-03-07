package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/response"
)

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
}

func NewAuthService(
	userRepository repositories.UserRepository,
	db *sql.DB,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *AuthServiceImpl) Login(
	ctx context.Context,
	param params.Auth,
) (response.LoginResponse, error) {
	user, err := service.UserRepository.FindByCredentials(ctx, service.DB, param.User)
	if err != nil {
		return response.LoginResponse{}, err
	}

	err = helper.CompareHashWithPassword(param.User.Password, user.Password)
	if err != nil {
		return response.LoginResponse{}, err
	}

	token, err := helper.GenerateJwtToken(user)
	if err != nil {
		return response.LoginResponse{}, err
	}

	return response.LoginResponse{
		Token: token,
	}, nil
}
