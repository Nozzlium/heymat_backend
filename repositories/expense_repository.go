package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/results"
)

type ExpenseRepository struct{}

func (repository *ExpenseRepository) Create(
	ctx context.Context,
	db *sql.DB,
	entity entities.Expense,
) (entities.Expense, error) {
	query := `
    insert into expense(
      title,
      amount,
      notes,
      user_id,
      budget_id,
      created_at,
      updated_at
    )
    values
    ($1, $2, $3, $4, $5, $6);
  `
	var insertedId uint64
	currentTime := time.Now()
	err := db.QueryRowContext(
		ctx,
		query,
		entity.Title,
		entity.Amount,
		entity.Notes,
		entity.UserID,
		entity.BudgetID,
		currentTime,
	).Scan(&insertedId)
	if err != nil {
		return entities.Expense{}, err
	}

	entity.ID = insertedId
	return entity, nil
}

func (repository *ExpenseRepository) Get(
	ctx context.Context,
	db *sql.DB,
	param params.Expense,
) ([]results.Expense, error) {
	query := `
    select 
      id,
      title,
      amount,
      notes,
      user_account.id,
      user_account.username,
      budget_id,
      created_at,
      updated_at
    from expense 
      join user_account on expense.user_id = user.id 
    where expense.budget_id = $1
    offset $2
    limit $3
    order by expense.created_at $4
  `
	rows, err := db.QueryContext(
		ctx,
		query,
		param.Expense.BudgetID,
		(param.PageNo-1)*param.PageSize,
		param.PageSize,
		param.Order,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]results.Expense, 0, param.PageSize)
	var notes sql.NullString
	for rows.Next() {
		result := results.Expense{}
		err := rows.Scan(
			&result.ID,
			&result.Title,
			&notes,
			&result.UserID,
			&result.Username,
			&result.BudgetID,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, result)
	}

	return res, nil
}
