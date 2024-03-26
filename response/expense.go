package response

import "time"

type ExpenseResponse struct {
	ID              uint64    `json:"id"`
	Title           string    `json:"title"`
	Amount          uint64    `json:"amount"`
	AmountString    string    `json:"amountString"`
	Notes           string    `json:"notes"`
	BudgetID        uint64    `json:"budgetId"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedAtString string    `json:"createdAtString"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedAtString string    `json:"updatedAtString"`
	UserResponse
}

type ExpensesResponse struct {
	PageNo   uint32            `json:"pageNo"`
	PageSize uint32            `json:"pageSize"`
	Expenses []ExpenseResponse `json:"expense"`
}
