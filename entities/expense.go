package entities

import "time"

type Expense struct {
	ID        uint64
	Title     string
	Amount    uint64
	Notes     string
	UserID    uint64
	BudgerID  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
