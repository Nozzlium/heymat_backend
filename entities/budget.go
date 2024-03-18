package entities

import "time"

type Budget struct {
	ID        uint32
	UserID    uint32
	Amount    uint64
	TimeCode  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
