package main

import (
	"fmt"
	"restApi/internal/config"
	"restApi/internal/logger"
)

func main() {
	// TODO - init config (viper, cleanenv)
	cfg := config.MustLoadConfig()

	// TODO - init logger (slog)
	log := logger.SetupLogger(&cfg.Logger)
	log.Info("Starting server")
	// TODO - init storage (postgres, redis);
	// TODO - init router
	// TODO - init server
}
