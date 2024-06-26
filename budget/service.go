package budget

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

var DB *sql.DB

var ErrUserNotFound = errors.New(
	"user tidak ada",
)

type BudgetPlanParam struct {
	PageNo     uint32
	PageSize   uint32
	BudgetPlan BudgetPlan
}

func createBudgetPlan(
	ctx context.Context,
	budgetPlan BudgetPlan,
) (BudgetPlanResponse, error) {
	currentTime := time.Now()
	budgetPlan.CreatedAt = currentTime
	budgetPlan.UpdatedAt = currentTime
	saved, err := save(
		ctx,
		DB,
		budgetPlan,
	)
	if err != nil {
		if driverErr, ok := err.(*pq.Error); ok {
			if driverErr.Code.Name() == "foreign_key_violation" {
				return BudgetPlanResponse{}, ErrUserNotFound
			}
		}
		return BudgetPlanResponse{}, err
	}

	return findBudgetPlanItemById(
		ctx,
		DB,
		saved.ID,
	)
}

func getBudgetPlanItemList(
	ctx context.Context,
	param BudgetPlanParam,
) (BudgetPlanListResponse, error) {
	resp, err := findAllBudgetPlanItems(
		ctx,
		DB,
		param,
	)
	if err != nil {
		return BudgetPlanListResponse{}, err
	}

	return resp, nil
}

func getBudgetPlanById(
	ctx context.Context,
	budgetPlan BudgetPlan,
) (BudgetPlanResponse, error) {
	resp, err := findBudgetPlanItemById(
		ctx,
		DB,
		budgetPlan.ID,
	)
	if err != nil {
		return BudgetPlanResponse{}, err
	}

	if resp.Private &&
		resp.User.ID != budgetPlan.UserID {
		return BudgetPlanResponse{}, errors.New(
			"not found",
		)
	}

	return resp, nil
}
