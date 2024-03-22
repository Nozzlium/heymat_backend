package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/results"
)

type BudgetPlanRepositoryImpl struct{}

func NewBudgetRepository() *BudgetPlanRepositoryImpl {
	return &BudgetPlanRepositoryImpl{}
}

func (repository *BudgetPlanRepositoryImpl) Create(
	ctx context.Context,
	DB *sql.DB,
	entity entities.BudgetPlan,
) (entities.BudgetPlan, error) {
	query := `
    insert into budget(
      user_id,
      amount,
      time_code,
      created_at,
      updated_at
    ) values (
      $1, $2, $3, $4, $4
    ) returning id;
  `
	var insertedId uint32
	err := DB.QueryRowContext(ctx, query, entity.UserID, entity.Amount, entity.TimeCode, entity.CreatedAt).
		Scan(&insertedId)
	if err != nil {
		return entities.BudgetPlan{}, err
	}

	entity.ID = insertedId
	return entity, nil
}

func (repository *BudgetPlanRepositoryImpl) Get(
	ctx context.Context,
	DB *sql.DB,
	param params.BudgetPlan,
) ([]results.BudgetPlanBalanceResult, error) {
	query := `
    select 
      budget.id,
      budget.time_code as date,
      budget.amount,
      sum(report_entries.amount) as balance 
    from budget 
      left join report_entries on budget.id = report_entries.budget_id
    where date_part('year', budget.time_code) = $1 
      and budget.user_id = $2
    limit $3
    offset $4
    group by budget.id order by date asc;
  `
	rows, err := DB.QueryContext(
		ctx,
		query,
		param.BudgetPlan.CreatedAt,
		param.BudgetPlan.UserID,
		param.PageSize,
		(param.PageNo-1)*param.PageSize,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]results.BudgetPlanBalanceResult, 0, param.PageSize)
	for rows.Next() {
		balance := results.BudgetPlanBalanceResult{}
		var expense sql.NullInt64
		err = rows.Scan(&balance.ID, &balance.Date, &balance.Budget, &expense)
		if err != nil {
			return nil, err
		}
		if expense.Valid {
			balance.Expense = uint64(expense.Int64)
		} else {
			balance.Expense = 0
		}

		res = append(res, balance)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(res)

	return res, nil
}

func (repository *BudgetPlanRepositoryImpl) GetById(
	ctx context.Context,
	DB *sql.DB,
	param params.BudgetPlan,
) (results.BudgetPlanBalanceResult, error) {
	// TODO
	return results.BudgetPlanBalanceResult{}, nil
}
