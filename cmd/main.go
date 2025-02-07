package main

import (
	"os"
	"restApi/internal/config"
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
	// TODO - init server
}
