package response

import "time"

type ReportEntry struct {
	Title           string    `json:"title"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedAtString string    `json:"createdAtString"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
