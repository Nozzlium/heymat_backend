package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/results"
)

type RecordEntryRepositoryImpl struct{}

func NewRecordEntryRepository() *RecordEntryRepositoryImpl {
	return &RecordEntryRepositoryImpl{}
}

func (repository *RecordEntryRepositoryImpl) Create(
	ctx context.Context,
	DB *sql.DB,
	entity entities.ReportEntry,
) (entities.ReportEntry, error) {
	query := `
    insert into report_entries
    (user_id, title, amount, time_code, created_at, updated_at)
    values ($1, $2, $3, $4, $5, $6)
    returning id;
  `
	now := time.Now()
	var insertedId uint32
	err := DB.QueryRowContext(ctx, query, entity.UserID, entity.Title, entity.Amount, helper.TruncateToMonth(now), now, now).
		Scan(&insertedId)
	if err != nil {
		return entity, err
	}

	entity.ID = insertedId
	return entity, err
}

func (repository *RecordEntryRepositoryImpl) GetYearly(
	ctx context.Context,
	DB *sql.DB,
	param params.Balance,
) (map[uint]results.MonthlyBalance, error) {
	fmt.Println(param)
	query := `
    select 
      budget.time_code as date,
      budget.amount,
      sum(report_entries.amount) as balance 
    from budget 
      left join report_entries on budget.id = report_entries.budget_id
    where date_part('year', budget.time_code) = $1 
      and budget.user_id = $2
    group by budget.id order by date asc;
  `
	rows, err := DB.QueryContext(ctx, query, param.Date.Year(), param.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[uint]results.MonthlyBalance)
	for rows.Next() {
		balance := results.MonthlyBalance{}
		var expense sql.NullInt64
		err = rows.Scan(&balance.Date, &balance.Budget, &expense)
		if err != nil {
			return nil, err
		}
		if expense.Valid {
			balance.Expense = uint64(expense.Int64)
		} else {
			balance.Expense = 0
		}

		res[uint(balance.Date.Month())] = balance
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(res)

	return res, nil
}

func (repository *RecordEntryRepositoryImpl) GetTotalOfMonth(
	ctx context.Context,
	DB *sql.DB,
	time time.Time,
) (results.MonthlyBalance, error) {
	query := `
    select 
      report_entries.time_code as date,
      budget.amount,
      sum(report_entries.amount) as balance 
    from budget 
      left join report_entries on report_entries.time_code = budget.time_code
    where date = date_trunc('month', $1)
      group by record_entries.time_code limit 1;
  `
	res := results.MonthlyBalance{}
	err := DB.QueryRowContext(ctx, query, time).Scan(&res.Date, &res.Expense)
	return res, err
}

func (repository *RecordEntryRepositoryImpl) GetByMonth(
	ctx context.Context,
	DB *sql.DB,
	param params.ReportEntry,
) ([]entities.ReportEntry, error) {
	query := `
    select
      id,
      title,
      amount,
      user_id,
      created_at,
      updated_at
    from report_entries
    where
      date_trunc('month', created_at) = date_trunc('month', $1) and
      deleted_at is null
    limit $2
    offset $3
  `

	reports := make([]entities.ReportEntry, param.PageSize)
	rows, err := DB.QueryContext(ctx,
		query,
		param.RecordEntry.CreatedAt,
		param.PageSize,
		((param.PageNo - 1) * param.PageSize),
	)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		res := entities.ReportEntry{}
		err := rows.Scan(
			&res.ID,
			&res.Title,
			&res.Amount,
			&res.UserID,
			&res.CreatedAt,
			&res.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return reports, err
}

func (repository *RecordEntryRepositoryImpl) GetById(
	ctx context.Context,
	DB *sql.DB,
	id uint32,
) (entities.ReportEntry, error) {
	query := `
  select 
    id,
    title,
    amount,
    user_id,
    created_at,
    updated_at
  from report_entries
  where id = $1
  limit 1
  `

	res := entities.ReportEntry{}
	err := DB.QueryRowContext(ctx, query, id).Scan(&res.ID,
		&res.Title,
		&res.Amount,
		&res.UserID,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return res, err
}

func (repository *RecordEntryRepositoryImpl) Edit(
	ctx context.Context,
	DB *sql.DB,
	entity entities.ReportEntry,
) (entities.ReportEntry, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *RecordEntryRepositoryImpl) Delete(
	ctx context.Context,
	DB *sql.DB,
	id uint32,
) (entities.ReportEntry, error) {
	panic("not implemented") // TODO: Implement
}
