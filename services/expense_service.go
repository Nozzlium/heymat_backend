package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/response"
)

type ExpenseService struct {
	ExpenseRepository repositories.ExpenseRepository
	DB                *sql.DB
}

func (service *ExpenseService) Create(
	ctx context.Context,
	entity entities.Expense,
) (response.ExpenseResponse, error) {
	ent, err := service.ExpenseRepository.Create(
		ctx, service.DB, entity)

	return helper.ExpenseEntityToExpenseResponse(ent), err
}
