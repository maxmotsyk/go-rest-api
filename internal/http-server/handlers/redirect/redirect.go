package redirect

import (
	"log/slog"
	"net/http"
)

type GETurl interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGeter GETurl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo get alias from request

		// todo get url by alias

		// todo redirect to url

	}
}
