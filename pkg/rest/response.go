package rest

import "github.com/go-playground/validator/v10"

type QueryInfo struct {
	QueryID          int64   `json:"query_id"`
	Query            string  `json:"query"`
	MaxExecutionTime float64 `json:"execution_time"`
}

type FindQueriesResponse []QueryInfo

type BadRequestResponse struct {
	Msg          string                 `json:"msg"`
	Error        string                 `json:"error,omitempty"`
	FieldsErrors map[string]interface{} `json:"fields,omitempty"`
}

func NewBadRequestResponse(msg string, err error) BadRequestResponse {
	resp := BadRequestResponse{
		Msg: msg,
	}

	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		resp.Error = err.Error()
		return resp
	}

	resp.FieldsErrors = make(map[string]interface{})
	for _, err := range errors {
		resp.FieldsErrors[err.Field()] = err.Error()
	}

	return resp
}

type InternalServerErrorResponse struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}

func NewInternalServerErrorResponse(msg string, error error) InternalServerErrorResponse {
	return InternalServerErrorResponse{Msg: msg, Error: error.Error()}
}
