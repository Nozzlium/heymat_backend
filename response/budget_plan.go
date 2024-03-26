package response

import "time"

type BudgetPlanResponse struct {
	ID      uint64 `json:"id"`
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}

type BudgetPlanBalanceResponse struct {
	ID              uint64       `json:"id"`
	Title           string       `json:"title"`
	User            UserResponse `json:"user"`
	CreatedAt       time.Time    `json:"date"`
	CreatedAtString string       `json:"dateString"`
	Balance         int64        `json:"balance"`
	BalanceString   string       `json:"balanceString"`
	Budget          uint64       `json:"budget"`
	BudgetString    string       `json:"budgetString"`
	Expense         uint64       `json:"expense"`
	ExpenseString   string       `json:"expenseString"`
}

type PaginatedBudgetPlanBalance struct {
	PageNo      uint32                      `json:"pageNo"`
	PageSize    uint32                      `json:"pageSize"`
	BudgetPlans []BudgetPlanBalanceResponse `json:"budgetPlans"`
}
