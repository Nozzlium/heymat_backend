package entities

import "time"

type BudgetPlan struct {
	ID        uint64
	UserID    uint64
	Amount    uint64
	Private   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
