package rest

import "github.com/go-playground/validator/v10"

// BadRequestResponse represents bad request body.
type BadRequestResponse struct {
	Msg          string                 `json:"msg"`
	Error        string                 `json:"error,omitempty"`
	FieldsErrors map[string]interface{} `json:"fields,omitempty"`
}

// NewBadRequestResponse constructor for BadRequestResponse
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

// InternalServerErrorResponse represents internal server body
type InternalServerErrorResponse struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}

// NewInternalServerErrorResponse constructor for InternalServerErrorResponse.
func NewInternalServerErrorResponse(msg string, error error) InternalServerErrorResponse {
	return InternalServerErrorResponse{Msg: msg, Error: error.Error()}
}
