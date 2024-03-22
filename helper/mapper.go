package helper

import (
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/response"
	"github.com/nozzlium/heymat_backend/results"
)

func UserEntityToResponseMapper(entity entities.User) response.UserResponse {
	return response.UserResponse{
		ID:               entity.ID,
		Username:         entity.Username,
		Email:            entity.Email,
		IsEmailConfirmed: entity.IsEmailConfirmed,
	}
}

func BudgetEntityToBudgetResponse(entity entities.BudgetPlan) {
}

func BudgetResultToBudgetRepsonseMapper(
	result results.BudgetPlanBalanceResult,
) response.BudgetPlanBalanceResponse {
	dateString := GetIdDateStringMonth(result.Date)
	budget := result.Budget
	expense := result.Expense
	balance := int64(budget - expense)
	return response.BudgetPlanBalanceResponse{
		Budget:        budget,
		BudgetString:  IntToCurrency(int64(budget)),
		Expense:       expense,
		ExpenseString: IntToCurrency(int64(expense)),
		Balance:       balance,
		BalanceString: IntToCurrency(balance),
		Date:          result.Date,
		DateString:    dateString,
	}
}

func BudgetResultsToBudgetResponses(
	res []results.BudgetPlanBalanceResult,
) []response.BudgetPlanBalanceResponse {
	resp := make([]response.BudgetPlanBalanceResponse, 0, len(res))
	for _, result := range res {
		resp = append(resp, BudgetResultToBudgetRepsonseMapper(result))
	}
	return resp
}
