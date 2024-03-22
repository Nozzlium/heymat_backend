package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
    insert into budget_plan(
      user_id,
      amount, 
      title,
      private,
      created_at,
      updated_at
    ) values (
      $1, $2, $3, $4, $5, $5
    ) returning id;
  `
	var insertedId uint64
	currentTime := time.Now()
	err := DB.QueryRowContext(ctx, query, entity.UserID, entity.Amount, entity.Title, entity.Private, currentTime).
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
      budget_plan.id,
      budget_plan.created_at as date,
      budget_plan.title,
      budget_plan.amount,
      budget_plan.private,
      sum(expense.amount) as expense,
      user_account.id as user_id,
      user_account.username as username,
      user_account.email as email
    from budget_plan 
      left join expense on budget_plan.id = expense.budget_id
      join user_account on user_account.id = budget_plan.user_id
    where budget_plan.user_id = $1 and budget_plan.deleted_at is null 
    group by (budget_plan.id, user_account.id) order by date desc 
    limit $2
    offset $3;
  `
	rows, err := DB.QueryContext(
		ctx,
		query,
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
		err = rows.Scan(
			&balance.ID,
			&balance.Date,
			&balance.Title,
			&balance.Amount,
			&balance.Private,
			&expense,
			&balance.UserID,
			&balance.Username,
			&balance.Email,
		)
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
