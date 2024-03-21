package response

import "time"

type MonthlyBalance struct {
	Date          time.Time `json:"date"`
	DateString    string    `json:"dateString"`
	Balance       uint64    `json:"balance"`
	BalanceString string    `json:"balanceString"`
	Budget        uint64    `json:"budget"`
	BudgetString  string    `json:"budgetString"`
	Expense       uint64    `json:"expense"`
	ExpenseString string    `json:"expenseString"`
}

type MonthlyBalances struct {
	MonthlyBalances []MonthlyBalance `json:"monthlyBalances"`
}
