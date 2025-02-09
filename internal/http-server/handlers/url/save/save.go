package save

import (
	"log/slog"
	"net/http"
	resp "restApi/internal/api/response"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
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
		// Parse regest body
		req := &Request{}
		err := render.DecodeJSON(r.Body, req)

		if err != nil {
			log.Error("error decoding request body", err.Error())
			render.JSON(w, r, resp.Error("error decoding request body"))
			return
		}

		log.Info("saving url", slog.String("url", req.URL), slog.String("alias", req.Alias))

		// Validate request
		if err := validator.New().Struct(req); err != nil {
			log.Error("error validating request", err.Error())
			render.JSON(w, r, resp.Error("error validating request"))
			return
		}
	}
}
