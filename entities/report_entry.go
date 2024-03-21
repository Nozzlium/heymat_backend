package entities

import "time"

type ReportEntry struct {
	ID        uint32
	Title     string
	Amount    int64
	Notes     string
	UserID    uint32
	TimeCode  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
