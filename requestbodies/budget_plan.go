package requestbodies

type BudgetPlan struct {
	Title   string `json:"title"`
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}
