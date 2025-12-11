package httperror

type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
