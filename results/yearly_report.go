package results

import "time"

type MonthlyBalance struct {
	Date    time.Time
	Budget  uint64
	Expense uint64
}
