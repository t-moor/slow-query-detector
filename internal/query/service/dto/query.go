package dto

// FindQueriesInput represents input data
type FindQueriesInput struct {
	Command string
	Sort    string
	Page    int
	PerPage int
}

// QueryInfo represents data for particular query
type QueryInfo struct {
	QueryID       int64
	Query         string
	ExecutionTime float64
}

// FindQueriesOutput represents output
type FindQueriesOutput []*QueryInfo
