package repository

type FindQueriesCriteria struct {
	Command string
	Sort    string
	Page    int
	PerPage int
}

type FindQueriesResultRow struct {
	QueryID          int64   `gorm:"column:queryid"`
	Query            string  `gorm:"column:query"`
	MaxExecutionTime float64 `gorm:"column:max_exec_time"`
}

func (g FindQueriesResultRow) TableName() string {
	return "pg_stat_statements"
}

type FindQueriesResults []*FindQueriesResultRow
