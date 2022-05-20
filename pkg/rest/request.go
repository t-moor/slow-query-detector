package rest

type FindQueriesRequest struct {
	Command string `validate:"omitempty,oneof=select insert update delete"`
	Sort    string `validate:"omitempty,oneof=slow-to-fast fast-to-slow"`
	Page    int    `validate:"gt=0"`
	PerPage int    `validate:"gt=0"`
}
