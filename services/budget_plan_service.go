package services

import (
	"context"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/response"
)

type BudgetService interface {
	Create(
		ctx context.Context,
		entity entities.BudgetPlan,
	) (response.BudgetPlanBalanceResponse, error)
	Get(ctx context.Context, param params.BudgetPlan) (response.BudgetPlanBalanceResponse, error)
}
