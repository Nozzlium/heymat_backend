package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/response"
)

type BudgetService interface {
	Create(ctx context.Context, Db *sql.DB, entity entities.Budget) (response.BudgetRepsonse, error)
}
