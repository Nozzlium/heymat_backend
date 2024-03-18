package requestbodies

import "time"

type Budget struct {
	Amount uint64    `json:"amount"`
	Date   time.Time `json:"date"`
}
