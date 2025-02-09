package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "ok"
	StatusError = "error"
)

// OK returns response with status OK
func OK() *Response {
	return &Response{Status: StatusOK}
}

// Error returns response with error message
func Error(msg string) *Response {
	return &Response{Status: StatusError, Error: msg}
}
