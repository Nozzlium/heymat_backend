package entities

import "time"

type ReportEntry struct {
	ID        uint32
	Title     string
	Amount    int64
	Notes     string
	UserID    uint32
	Year      uint16
	Month     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
