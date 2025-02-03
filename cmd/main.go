package main

import (
	"os"
	"restApi/internal/config"
	"restApi/internal/lib/logger/sl"
	"restApi/internal/logger"
	"restApi/internal/storage/sqlite"
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

	// TODO - init router

	// TODO - init server
}
