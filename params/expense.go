package params

import "github.com/nozzlium/heymat_backend/entities"

type Expense struct {
	PageNo   uint32
	PageSize uint32
	Order    string
	Expense  entities.Expense
}
