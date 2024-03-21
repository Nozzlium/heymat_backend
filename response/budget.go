package response

import "time"

type BudgetRepsonse struct {
	Amount     uint64    `json:"amount"`
	Date       time.Time `json:"date"`
	DateString string    `json:"dateString"`
}
