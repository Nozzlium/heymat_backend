package helper

import (
	"github.com/nozzlium/heymat_backend/entities"
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

func ReportEntryEntityToResponseMapper(entity entities.ReportEntry) response.ReportEntryResponse {
	// TODO implement date Indonesia
	return response.ReportEntryResponse{
		Title:     entity.Title,
		Notes:     entity.Notes,
		CreatedAt: entity.CreatedAt,
	}
}
