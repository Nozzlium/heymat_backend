package budget

import (
	"context"
	"database/sql"
	"log"
)

func save(
	ctx context.Context,
	DB *sql.DB,
	entity BudgetPlan,
) (BudgetPlan, error) {
	query := `
    insert into budget_plan(
      user_id,
      amount, 
      title,
      private,
      created_at,
      updated_at 
    ) values (
      $1, $2, $3, $4, $5, $6
    ) returning id;
  `
	var insertedId uint64
	err := DB.QueryRowContext(
		ctx,
		query,
		entity.UserID,
		entity.Amount,
		entity.Title,
		entity.Private,
		entity.CreatedAt,
		entity.UpdatedAt,
	).Scan(&insertedId)
	if err != nil {
		log.Println(err)
		return BudgetPlan{}, err
	}

	entity.ID = insertedId
	return entity, nil
}

func findById(
	ctx context.Context,
	DB *sql.DB,
	id uint64,
) (BudgetPlan, error) {
	query := ` 
    select 
      id,
      user_id,
      title,
      amount,
      private,
      created_at,
      updated_at,
      deleted_at
    from budget_plan
    where id = $1
  `
	budgetPlan := BudgetPlan{}
	err := DB.QueryRowContext(ctx, query, id).
		Scan(
			&budgetPlan.ID,
			&budgetPlan.UserID,
			&budgetPlan.Title,
			&budgetPlan.Amount,
			&budgetPlan.Private,
			&budgetPlan.CreatedAt,
			&budgetPlan.UpdatedAt,
			&budgetPlan.DeletedAt,
		)
	if err != nil {
		log.Println(err)
		return BudgetPlan{}, err
	}

	return budgetPlan, err
}
