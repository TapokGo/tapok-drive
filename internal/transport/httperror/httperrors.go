// Package httperror provides HTTPError methods
package httperror

import (
	"net/http"
)

// HTTPError is a model of http error
type HTTPError struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
}

func InvalidRequest(message string) HTTPError {
	return HTTPError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func InternalError(message string) HTTPError {
	return HTTPError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func ConflictError(message string) HTTPError {
	return HTTPError{
		Message: message,
		Code:    http.StatusConflict,
	}
}
