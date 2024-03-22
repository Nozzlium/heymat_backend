package repositories

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/results"
)

type BudgetPlanRepository interface {
	Create(ctx context.Context, DB *sql.DB, entity entities.BudgetPlan) (entities.BudgetPlan, error)
	Get(
		ctx context.Context,
		DB *sql.DB,
		param params.BudgetPlan,
	) ([]results.BudgetPlanBalanceResult, error)
	GetById(
		ctx context.Context,
		DB *sql.DB,
		param params.BudgetPlan,
	) (results.BudgetPlanBalanceResult, error)
}
