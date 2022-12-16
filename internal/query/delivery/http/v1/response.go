package v1

const internalServerError = "something went wrong"

type Payload struct {
	ID                int64   `json:"id"`
	Statement         string  `json:"statement"`
	MaxExecutionTime  float64 `json:"max_exec_time"`
	MeanExecutionTime float64 `json:"mean_exec_time"`
}

type Response struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	PageCount  int       `json:"page_count"`
	TotalCount int       `json:"total_count"`
	Data       []Payload `json:"data"`
}
