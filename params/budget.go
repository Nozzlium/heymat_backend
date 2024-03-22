package params

import (
	"github.com/nozzlium/heymat_backend/entities"
)

type BudgetPlan struct {
	PageNo     uint32
	PageSize   uint32
	BudgetPlan entities.BudgetPlan
}
