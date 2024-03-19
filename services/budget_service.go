package services

import (
	"context"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/response"
)

type BudgetService interface {
	Create(ctx context.Context, entity entities.Budget) (response.BudgetRepsonse, error)
}
