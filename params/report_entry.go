package params

import "github.com/nozzlium/heymat_backend/entities"

type ReportEntry struct {
	PageNo      uint
	PageSize    uint
	RecordEntry entities.ReportEntry
}
