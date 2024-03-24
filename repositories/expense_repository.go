package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
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
		entity.BudgerID,
		currentTime,
	).Scan(&insertedId)
	if err != nil {
		return entities.Expense{}, err
	}

	entity.ID = insertedId
	return entity, nil
}
