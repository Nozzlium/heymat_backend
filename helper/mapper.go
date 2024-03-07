package helper

import (
	"github.com/nozzlium/heymat_backend/data/entities"
	"github.com/nozzlium/heymat_backend/response"
)

func UserEntityToResponseMapper(entity entities.User) response.UserResponse {
	return response.UserResponse{
		ID:               entity.ID,
		Username:         entity.Username,
		Email:            entity.Email,
		IsEmailConfirmed: entity.IsEmailConfirmed,
	}
}
