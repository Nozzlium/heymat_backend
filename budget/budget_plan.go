package budget

import "time"

type BudgetPlan struct {
	ID        uint64
	UserID    uint64
	Title     string
	Amount    uint64
	Private   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
