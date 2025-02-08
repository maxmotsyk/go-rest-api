package save

import (
	"log/slog"
	"net/http"
	resp "restApi/internal/api/response"
)

// Struct for request
type Request struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

// Struct for response
type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

// Interface for URLSaver
type URLSaver interface {
	Save(urlToSave string, alias string) error
}

// New function for save handler
func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
