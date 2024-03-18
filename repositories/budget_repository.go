package repositories

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
)

type BudgetRepository interface {
	Create(ctx context.Context, DB *sql.DB, entity entities.Budget) (entities.Budget, error)
}
