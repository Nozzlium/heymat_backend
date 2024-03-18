package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
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
) (response.ReportEntry, error) {
	res, err := service.ReportEntryRepository.Create(ctx, service.DB, entity)
	return helper.ReportEntryEntityToResponseMapper(res), err
}

func (service *ReportServiceImpl) GetByYear(
	ctx context.Context,
	param params.ReportEntry,
) ([]response.MonthlyBalance, error) {
	reports, err := service.ReportEntryRepository.GetYearly(
		ctx,
		service.DB,
		param.RecordEntry.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	year := param.RecordEntry.CreatedAt.Year()
	res := make([]response.MonthlyBalance, 12, 12)
	for i, temp := range res {
		monthInt := uint(i + 1)
		reprt, ok := reports[monthInt]
		if ok {
			temp.Budget = reprt.Budget
			temp.BudgetString = helper.IntToCurrency(temp.Budget)
			temp.Expense = reprt.Expense
			temp.ExpenseString = helper.IntToCurrency(temp.Expense)
			temp.Balance = temp.Budget - temp.Expense
			temp.BalanceString = helper.IntToCurrency(temp.Balance)
		}
		date := time.Date(year, time.Month(i), 1, 0, 0, 0, 0, time.UTC)
		temp.Date = date
		temp.DateString = helper.GetIdDateStringMonth(temp.Date)
	}
	return res, nil
}
