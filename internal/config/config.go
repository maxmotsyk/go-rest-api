package config

import "time"

type Config struct {
	Env        string     `yaml:"ENV"  env-default:"local"`
	HtttServer HTTPServer `yaml:"HTTP_SERVER" env-required:"true"`
	DataBase   DataBase   `yaml:"DATABASE"`
	Logger     Logger     `yaml:"LOGGER"`
}

type HTTPServer struct {
	Address     string        `yaml:"ADDRESS"  env-default:"local-host:8080"`
	ReadTimout  time.Duration `yaml:"READ_TIMEOUT"  env-default:"5s"`
	WriteTimout time.Duration `yaml:"WRITE_TIMEOUT"  env-default:"10s"`
	IDLETimeout time.Duration `yaml:"IDLE_TIMEOUT"  env-default:"60s"`
}

type DataBase struct {
	Host     string `yaml:"HOST"  env-default:"localhost"`
	Port     int    `yaml:"PORT"  env-default:"5432"`
	Name     string `yaml:"NAME"  env-default:"database_orders"`
	User     string `yaml:"USER"  env-default:"postgres" env-required:"true"`
	Password string `yaml:"PASS"  env-default:"password" env-required:"true"`
	SSLMode  string `yaml:"SSL_MODE"  env-default:"disable"`
}

type Logger struct {
	Level      string `yaml:"LEVEL"  env-default:"info"`
	FilePath   string `yaml:"FILE_PATH"  env-default:"./logs/app.log"`
	MaxSize    int    `yaml:"MAX_SIZE"  env-default:"30"`
	MaxBackups int    `yaml:"MAX_BACKUPS"  env-default:"3"`
	MaxAge     int    `yaml:"MAX_AGE"  env-default:"7"`
}
