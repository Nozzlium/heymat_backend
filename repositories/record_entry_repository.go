package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/results"
)

type RecordEntryRepository interface {
	Create(
		ctx context.Context,
		DB *sql.DB,
		entity entities.ReportEntry,
	) (entities.ReportEntry, error)
	GetYearly(ctx context.Context, DB *sql.DB, time time.Time) ([]results.YearlyReport, error)
	GetTotalOfMonth(ctx context.Context, DB *sql.DB, time time.Time) (results.YearlyReport, error)
	GetByMonth(ctx context.Context, DB *sql.DB, time time.Time) ([]entities.ReportEntry, error)
	GetById(ctx context.Context, DB *sql.DB, id uint32) (entities.ReportEntry, error)
	Edit(ctx context.Context, DB *sql.DB, entity entities.ReportEntry) (entities.ReportEntry, error)
	Delete(ctx context.Context, DB *sql.DB, id uint32) (entities.ReportEntry, error)
}