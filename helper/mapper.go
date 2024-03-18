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

func ReportEntryEntityToResponseMapper(entity entities.ReportEntry) response.ReportEntry {
	return response.ReportEntry{
		Title:           entity.Title,
		Notes:           entity.Notes,
		CreatedAt:       entity.CreatedAt,
		CreatedAtString: GetIdDateStringFull(entity.CreatedAt),
		UpdatedAt:       entity.UpdatedAt,
	}
}

func BudgetEntityToBudgetRepsonseMapper(entity entities.Budget) response.BudgetRepsonse {
	return response.BudgetRepsonse{
		Amount: entity.Amount,
		Date:   entity.CreatedAt,
	}
}
