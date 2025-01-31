package main

import (
	"fmt"
	"os"
	"restApi/internal/config"
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
	if err != nil {
		log.Error(fmt.Sprintf("Error initializing storage: %v", err))
		os.Exit(1)
		return
	}
	log.Info("Storage initialized")

	fmt.Println(storage)
	// TODO - init router
	// TODO - init server
}
