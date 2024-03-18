package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/response"
)

type BudgetServiceImpl struct {
	BudgetReporitory repositories.BudgetRepository
	DB               *sql.DB
}

func (service *BudgetServiceImpl) Create(
	ctx context.Context,
	Db *sql.DB,
	entity entities.Budget,
) (response.BudgetRepsonse, error) {
	res, err := service.BudgetReporitory.Create(ctx, service.DB, entity)
	return helper.BudgetEntityToBudgetRepsonseMapper(res), err
}
