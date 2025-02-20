package delete

import (
	"log/slog"
	"net/http"
	resp "restApi/internal/api/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type DELurl interface {
	DeleteByAlias(alias string) error
}

func New(log *slog.Logger, urlDelete DELurl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get alias from request
		alias := chi.URLParam(r, "alias")

		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("error validating request"))
			return
		}

		err := urlDelete.DeleteByAlias(alias)

		//delete url by alias
		if err != nil {
			log.Error("error deleting url by alias", slog.Any("alias", alias))
			render.JSON(w, r, resp.Error("error deleting url by alias"))
			return
		}

		log.Info("url deleted", slog.Any("alias", alias))
		render.JSON(w, r, resp.OK())
	}
}
