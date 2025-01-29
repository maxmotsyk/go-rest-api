package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string     `yaml:"env" env-default:"local"`
	HTTPServer HTTPServer `yaml:"http_server" env-required:"true"`
	DataBase   DataBase   `yaml:"db"`
	Logger     Logger     `yaml:"logger"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env-default:"5s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env-default:"10s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DataBase struct {
	Host        string `yaml:"host" env-default:"localhost"`
	Port        int    `yaml:"port" env-default:"5432"`
	Name        string `yaml:"name" env-default:"database_orders"`
	User        string `yaml:"user" env-default:"postgres"`
	Password    string `yaml:"password" env-default:"password"`
	StoragePath string `yaml:"storage_path"`
	SSLMode     string `yaml:"ssl_mode" env-default:"disable"`
}

type Logger struct {
	Level      string `yaml:"level" env-default:"info"`
	Envirment  string `yaml:"envirment" env-default:"dev"`
	FilePath   string `yaml:"file_path" env-default:"./logs/app.log"`
	MaxSize    int    `yaml:"max_size" env-default:"10"`
	MaxBackups int    `yaml:"max_backups" env-default:"3"`
	MaxAge     int    `yaml:"max_age" env-default:"30"`
}

func MustLoadConfig() *Config {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")

	fmt.Println(configPath)

	//check if config path is empty
	if configPath == "" {
		log.Fatal("config path is empty")
	}

	//check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	//load config file
	var config Config

	//load config file
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("config file read error: %s", err)
	}

	return &config

}
