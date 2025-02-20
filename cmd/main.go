package main

import (
	"net/http"
	"os"
	"restApi/internal/config"
	"restApi/internal/http-server/handlers/redirect"
	"restApi/internal/http-server/handlers/url/delete"
	"restApi/internal/http-server/handlers/url/save"
	mwLogger "restApi/internal/http-server/middleware/logger"
	"restApi/internal/lib/logger/sl"
	"restApi/internal/logger"
	"restApi/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// TODO - init config (viper, cleanenv)
	cfg := config.MustLoadConfig()

	// TODO - init logger (slog)
	log := logger.SetupLogger(&cfg.Logger)
	log.Info("Starting server")
	// TODO - init storage (postgres, redis);
	storage, err := sqlite.NewStorage(cfg.DataBase.StoragePath)

	defer storage.Close()

	if err != nil {
		log.Error("Error initializing storage:", sl.Err(err))
		os.Exit(1)
	}
	log.Info("Storage initialized")

	// TODO - init router (chi)
	router := chi.NewRouter()
	log.Info("Router initialized")

	// TODO - init middlewares
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url/save", save.New(log, storage))
	router.Get("/{alias}", redirect.New(log, storage))
	router.Delete("/url/{alias}", delete.New(log, storage))

	// TODO - init server
	serv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	// TODO - start server

	if err := serv.ListenAndServe(); err != nil {
		log.Error("Error starting server:", sl.Err(err))
		os.Exit(1)
	}

}
