package requestbodies

type Budget struct {
	Title   string `json:"title"`
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}
