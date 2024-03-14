package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/results"
)

type RecordEntryRepositoryImpl struct{}

func (repository *RecordEntryRepositoryImpl) Create(
	ctx context.Context,
	DB *sql.DB,
	entity entities.ReportEntry,
) (entities.ReportEntry, error) {
	query := `
    insert into report_entries
    (user_id, title, amount, created_at, updated_at)
    values ($1, $2, $3, $4, $5)
    returning id;
  `
	now := time.Now()
	var insertedId uint32
	err := DB.QueryRowContext(ctx, query, entity.UserID, entity.Title, entity.Amount, now, now).
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
	time time.Time,
) (map[uint]results.YearlyReport, error) {
	query := `
    select 
      date_trunc('month', created_at) as date,
      sum(value) as value
    from report_entries
    where date_part('year', created_at) = date_part('year', $1)
    group by date order by date;
  `
	rows, err := DB.QueryContext(ctx, query, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[uint]results.YearlyReport)
	for rows.Next() {
		mth := results.YearlyReport{}
		err = rows.Scan(&mth.Date, &mth.Value)
		if err != nil {
			return nil, err
		}

		res[uint(mth.Date.Month())] = mth
	}

	return res, nil
}

func (repository *RecordEntryRepositoryImpl) GetTotalOfMonth(
	ctx context.Context,
	DB *sql.DB,
	time time.Time,
) (results.YearlyReport, error) {
	query := `
    select 
      date_trunc('month', created_at) as date,
      sum(value) as value
    from record_entries
    where date = date_part('month', $1)  
  `
	res := results.YearlyReport{}
	err := DB.QueryRowContext(ctx, query, time).Scan(&res.Date, &res.Value)
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
