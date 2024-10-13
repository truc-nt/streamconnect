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

type ShopifyAppConfig struct {
	ClientID     string `env:"SHOPIFY_APP_CLIENT_ID"`
	ClientSecret string `env:"SHOPIFY_APP_CLIENT_SECRET"`
}

type VideoSdkConfig struct {
	Token string `env:"VIDEOSDK_TOKEN"`
}

type Config struct {
	Postgres   PostgresConfig
	Server     ServerConfig
	ShopifyApp ShopifyAppConfig
	VideoSdk   VideoSdkConfig
}

func NewConfig() *Config {
	var config Config

	//you can add a flag to check for env file if it is experimental environment
	//deployment cannot use env file
	if err := godotenv.Load(); err != nil {
		//ignore if can't find env file
	}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return &config
}
