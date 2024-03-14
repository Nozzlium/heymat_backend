package response

type Yearly struct {
	MonthInt uint8  `json:"monthInt"`
	Month    string `json:"month"`
	Sum      uint64 `json:"sum"`
}
