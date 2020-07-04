package errors

import (
	"net/http"
)

// RestErr struct
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// error const
const (
	errBadRequest     string = "bad_request"
	errNotFound       string = "not_found"
	errInternalServer string = "internal_server_error"
)

// NewBadRequestError func to send bad request error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   errBadRequest,
	}
}

// NewNotFoundError func to send not found error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   errNotFound,
	}
}

// NewInternalServerError func to send internal server error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   errInternalServer,
	}
}
