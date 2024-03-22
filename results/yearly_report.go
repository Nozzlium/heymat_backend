package results

import "time"

type BudgetPlanBalanceResult struct {
	ID      uint64
	Date    time.Time
	Budget  uint64
	Expense uint64
}
