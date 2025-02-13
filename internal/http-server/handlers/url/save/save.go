package save

import (
	"errors"
	"log/slog"
	"net/http"
	resp "restApi/internal/api/response"
	"restApi/internal/lib/random"
	"restApi/internal/storage"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

// Struct for request
type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

// Struct for response
type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

// Interface for URLSaver
type URLSaver interface {
	SaveURL(urlToSave string, alias string) error
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

		log.Info("reqest body deccode", slog.Any("reqest", req))

		// Validate request
		if err := validator.New().Struct(req); err != nil {
			log.Error("error validating request", err.Error())
			render.JSON(w, r, resp.Error("error validating request"))
			return
		}

		// Save URL
		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(6)
		}

		err = urlSaver.SaveURL(req.URL, alias)

		if errors.Is(err, storage.ErrorURLExists) {
			log.Error("error saving url", err.Error())
			render.JSON(w, r, resp.Error("alias already exists"))
			return
		} else if err != nil {
			log.Error("error saving url", err.Error())
			render.JSON(w, r, resp.Error("error saving url"))
			return
		}

		// Return response

		render.JSON(w, r, &Response{
			Response: *resp.OK(),
			Alias:    alias,
		})
	}

}
