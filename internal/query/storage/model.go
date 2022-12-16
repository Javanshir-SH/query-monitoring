package storage

const (
	QueryStatStatementsTableName = "pg_stat_statements"
)

type FilterOption struct {
	Command string
	Sort    string
	Limit   int
	Offset  int
}

type Query struct {
	QueryID           int64   `gorm:"column:queryid"`
	Query             string  `gorm:"column:query"`
	MaxExecutionTime  float64 `gorm:"column:max_exec_time"`
	MeanExecutionTime float64 `gorm:"column:mean_exec_time"`
}

type ListOfQuery struct {
	TotalCount int
	Rows       []*Query
}

func (q Query) TableName() string {
	return QueryStatStatementsTableName
}
