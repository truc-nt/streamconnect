package configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DbName   string `env:"POSTGRES_DBNAME"`
	Port     int    `env:"POSTGRES_PORT"`
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST"`
	Port int    `env:"SERVER_PORT"`
}

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

func NewConfig() *Config {
	var config Config

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return &config
}
