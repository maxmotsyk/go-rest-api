package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "ok"
	StatusError = "error"
)

func OK() *Response {
	return &Response{Status: StatusOK}
}

func Error(err error) *Response {
	return &Response{Status: StatusError, Error: err.Error()}
}
