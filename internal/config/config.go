package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host    string `env:"CHANNEL_LEVEL_HOST" env-required:"true"`
	Port    int32  `env:"CHANNEL_LEVEL_PORT" env-required:"true"`
	BaseURL string `env:"TRANSPORT_LEVEL_URL" env-required:"true"`
}

func MustLoad() *Config {
	var cfg Config
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("unable to load .env file: %v", err)
		}
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}
	return &cfg
}
