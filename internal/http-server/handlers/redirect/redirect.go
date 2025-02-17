package redirect

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	resp "restApi/internal/api/response"
	"restApi/internal/storage"
)

type GETurl interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGeter GETurl) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// todo get alias from request
		alias := chi.URLParam(r, "alias")

		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("error validating request"))
			return
		}
		// todo get url by alias
		url, err := urlGeter.GetURL(alias)

		if errors.Is(err, storage.ErrorUrlsNotFound) {
			log.Error("url not found", slog.Any("alias", alias))
			render.JSON(w, r, resp.Error("url not found"))
			return
		}

		if err != nil {
			log.Error("error getting url by alias", slog.Any("alias", alias))
			render.JSON(w, r, resp.Error("error getting url by alias"))
			return
		}

		// todo redirect to url

		log.Info("redirect to url", slog.Any("url", url))
		http.Redirect(w, r, url, http.StatusFound)

	}

}
