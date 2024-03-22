package results

import "time"

type BudgetPlanBalanceResult struct {
	ID       uint64
	Date     time.Time
	Amount   uint64
	Private  bool
	Expense  uint64
	UserID   uint64
	Username string
	Email    string
}
