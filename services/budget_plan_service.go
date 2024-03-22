package services

import (
	"context"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/response"
)

type BudgetPlanService interface {
	Create(
		ctx context.Context,
		entity entities.BudgetPlan,
	) (response.BudgetPlanResponse, error)
	Get(ctx context.Context, param params.BudgetPlan) (response.BudgetPlanBalanceResponses, error)
}
