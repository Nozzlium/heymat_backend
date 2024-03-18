package repositories

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
)

type BudgetRepositoryImpl struct{}

func (repository *BudgetRepositoryImpl) Create(
	ctx context.Context,
	DB *sql.DB,
	entity entities.Budget,
) (entities.Budget, error) {
	query := `
    insert into budget(
      user_id,
      amount,
      time_code,
      created_at,
      updated_at
    ) values (
      $1, $2, date_trunc('month', $3), $3, $3
    ) returning id;
  `
	var insertedId uint32
	err := DB.QueryRowContext(ctx, query, entity.UserID, entity.Amount, entity.CreatedAt).
		Scan(&insertedId)
	if err != nil {
		return entities.Budget{}, err
	}

	entity.ID = insertedId
	return entity, nil
}
