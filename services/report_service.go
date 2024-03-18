package services

import (
	"context"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/response"
)

type ReportService interface {
	Create(ctx context.Context, entity entities.ReportEntry) (response.ReportEntry, error)
	GetByYear(ctx context.Context, param params.ReportEntry) ([]response.MonthlyBalance, error)
}
