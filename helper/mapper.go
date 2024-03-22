package helper

import (
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/response"
	"github.com/nozzlium/heymat_backend/results"
)

func UserEntityToResponseMapper(entity entities.UserAccount) response.UserResponse {
	return response.UserResponse{
		ID:               entity.ID,
		Username:         entity.Username,
		Email:            entity.Email,
		IsEmailConfirmed: entity.IsEmailConfirmed,
	}
}

func BudgetEntityToBudgetResponse(entity entities.BudgetPlan) response.BudgetPlanResponse {
	return response.BudgetPlanResponse{
		ID:      entity.ID,
		Amount:  entity.Amount,
		Private: entity.Private,
	}
}

func BudgetResultToBudgetRepsonseMapper(
	result results.BudgetPlanBalanceResult,
) response.BudgetPlanBalanceResponse {
	dateString := GetIdTimeStringFull(result.Date)
	budget := result.Amount
	expense := result.Expense
	balance := int64(budget - expense)
	return response.BudgetPlanBalanceResponse{
		ID:    result.ID,
		Title: result.Title,
		User: response.UserResponse{
			ID:       result.UserID,
			Username: result.Username,
			Email:    result.Email,
		},
		Budget:          budget,
		BudgetString:    IntToCurrency(int64(budget)),
		Expense:         expense,
		ExpenseString:   IntToCurrency(int64(expense)),
		Balance:         balance,
		BalanceString:   IntToCurrency(balance),
		CreatedAt:       result.Date,
		CreatedAtString: dateString,
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
