package services

import (
	"context"
	"database/sql"
	"fmt"
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

func NewReportService(
	reportEntryRepository repositories.RecordEntryRepository,
	DB *sql.DB,
) *ReportServiceImpl {
	return &ReportServiceImpl{
		ReportEntryRepository: reportEntryRepository,
		DB:                    DB,
	}
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
) (response.MonthlyBalances, error) {
	reports, err := service.ReportEntryRepository.GetYearly(
		ctx,
		service.DB,
		params.Balance{
			UserID: uint64(param.RecordEntry.UserID),
			Date:   param.RecordEntry.CreatedAt,
		},
	)
	if err != nil {
		return response.MonthlyBalances{}, err
	}

	year := param.RecordEntry.CreatedAt.Year()
	res := make([]response.MonthlyBalance, 12, 12)
	for i, temp := range res {
		monthInt := uint(i + 1)
		reprt, ok := reports[monthInt]
		if ok {
			fmt.Println(reprt)
			temp.Budget = reprt.Budget
			temp.BudgetString = helper.IntToCurrency(temp.Budget)
			temp.Expense = reprt.Expense
			temp.ExpenseString = helper.IntToCurrency(temp.Expense)
			temp.Balance = temp.Budget - temp.Expense
			temp.BalanceString = helper.IntToCurrency(temp.Balance)
		}
		date := time.Date(
			year,
			time.Month(monthInt),
			1,
			0,
			0,
			0,
			0,
			time.FixedZone("GMT+7", 7*60*60),
		)
		temp.Date = date
		temp.DateString = helper.GetIdDateStringMonth(temp.Date)
	}
	return response.MonthlyBalances{
		MonthlyBalances: res,
	}, nil
}
