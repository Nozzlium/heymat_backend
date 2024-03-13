package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
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
) ([]results.YearlyReport, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *RecordEntryRepositoryImpl) GetTotalOfMonth(
	ctx context.Context,
	DB *sql.DB,
	time time.Time,
) (results.YearlyReport, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *RecordEntryRepositoryImpl) GetByMonth(
	ctx context.Context,
	DB *sql.DB,
	time time.Time,
) ([]entities.ReportEntry, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *RecordEntryRepositoryImpl) GetById(
	ctx context.Context,
	DB *sql.DB,
	id uint32,
) (entities.ReportEntry, error) {
	panic("not implemented") // TODO: Implement
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
