package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/response"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
}

func NewUserService(
	userRepository repositories.UserRepository,
	db *sql.DB,
) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Register(
	ctx context.Context,
	param params.User,
) (response.UserResponse, error) {
	user := param.User
	passwordHash, err := helper.HashPassword(user.Password)
	if err != nil {
		return response.UserResponse{}, err
	}

	user.Password = passwordHash
	result, err := service.UserRepository.Create(ctx, service.DB, user)
	return helper.UserEntityToResponseMapper(result), err
}
