package params

import "time"

type Budget struct {
	UserId uint64
	Amount uint64
	Date   time.Time
}
