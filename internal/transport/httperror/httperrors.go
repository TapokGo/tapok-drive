// Package httperror provides HTTPError methods
package httperror

// HTTPError is a model of http error
type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
