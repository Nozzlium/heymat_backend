package results

import "time"

type Expense struct {
	ID        uint64
	Title     string
	Amount    uint64
	Notes     string
	UserID    uint64
	Username  string
	BudgetID  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
