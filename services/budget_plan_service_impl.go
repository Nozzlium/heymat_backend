package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/response"
)

type BudgetPlanServiceImpl struct {
	BudgetReporitory repositories.BudgetPlanRepository
	DB               *sql.DB
}

func NewBudgetService(
	budgetRepsitory repositories.BudgetPlanRepository,
	DB *sql.DB,
) *BudgetPlanServiceImpl {
	return &BudgetPlanServiceImpl{
		BudgetReporitory: budgetRepsitory,
		DB:               DB,
	}
}

func (service *BudgetPlanServiceImpl) Create(
	ctx context.Context,
	entity entities.BudgetPlan,
) (response.BudgetPlanResponse, error) {
	res, err := service.BudgetReporitory.Create(ctx, service.DB, entity)
	return helper.BudgetEntityToBudgetResponse(res), err
}

func (service *BudgetPlanServiceImpl) Get(
	ctx context.Context,
	param params.BudgetPlan,
) (response.PaginatedBudgetPlanBalance, error) {
	res, err := service.BudgetReporitory.Get(ctx, service.DB, param)

	return response.PaginatedBudgetPlanBalance{
		PageNo:      param.PageNo,
		PageSize:    uint32(len(res)),
		BudgetPlans: helper.BudgetResultsToBudgetResponses(res),
	}, err
}
