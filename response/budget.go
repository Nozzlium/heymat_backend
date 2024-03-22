package response

import "time"

type BudgetPlanResponse struct{}

type BudgetPlanBalanceResponse struct {
	Date          time.Time `json:"date"`
	DateString    string    `json:"dateString"`
	Balance       int64     `json:"balance"`
	BalanceString string    `json:"balanceString"`
	Budget        uint64    `json:"budget"`
	BudgetString  string    `json:"budgetString"`
	Expense       uint64    `json:"expense"`
	ExpenseString string    `json:"expenseString"`
}

type BudgetPlanBalanceResponses struct {
	PageNo      uint32                      `json:"pageNo"`
	PageSize    uint32                      `json:"pageSize"`
	BudgetPlans []BudgetPlanBalanceResponse `json:"budgetPlans"`
}
