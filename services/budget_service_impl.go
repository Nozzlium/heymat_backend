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

type BudgetServiceImpl struct {
	BudgetReporitory repositories.BudgetPlanRepository
	DB               *sql.DB
}

func NewBudgetService(
	budgetRepsitory repositories.BudgetPlanRepository,
	DB *sql.DB,
) *BudgetServiceImpl {
	return &BudgetServiceImpl{
		BudgetReporitory: budgetRepsitory,
		DB:               DB,
	}
}

func (service *BudgetServiceImpl) Create(
	ctx context.Context,
	entity entities.BudgetPlan,
) (response.BudgetPlanBalanceResponse, error) {
	_, err := service.BudgetReporitory.Create(ctx, service.DB, entity)
	return response.BudgetPlanBalanceResponse{}, err
}

func (service *BudgetServiceImpl) Get(
	ctx context.Context,
	param params.BudgetPlan,
) (response.BudgetPlanBalanceResponses, error) {
	res, err := service.BudgetReporitory.Get(ctx, service.DB, param)

	return response.BudgetPlanBalanceResponses{
		PageNo:      param.PageNo,
		PageSize:    uint32(len(res)),
		BudgetPlans: helper.BudgetResultsToBudgetResponses(res),
	}, err
}
