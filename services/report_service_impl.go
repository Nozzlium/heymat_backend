package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/response"
)

type ReportServiceImpl struct {
	ReportEntryRepository repositories.RecordEntryRepository
	DB                    *sql.DB
}

func (service *ReportServiceImpl) Create(
	ctx context.Context,
	entity entities.ReportEntry,
) (response.ReportEntryResponse, error) {
	res, err := service.ReportEntryRepository.Create(ctx, service.DB, entity)
	if err != nil {
		return response.ReportEntryResponse{}, err
	}
	// TODO wait for date Indonesia
}

func (service *ReportServiceImpl) GetByYear(
	ctx context.Context,
	param params.ReportEntry,
) ([]response.Yearly, error) {
	reports, err := service.ReportEntryRepository.GetYearly(
		ctx,
		service.DB,
		param.RecordEntry.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	res := make([]response.Yearly, 12, 12)
	for i, temp := range res {
		monthInt := uint(i + 1)
		reprt, ok := reports[monthInt]
		if ok {
			temp.Sum = uint64(reprt.Value)
		}
		temp.MonthInt = uint8(monthInt)
		temp.Month = time.Month(monthInt).String()
	}
	return res, nil
}
