package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Adress string `env:"SERVER_ADDRESS" env-default:":8001"`
}

func NewServerConfig() (*ServerConfig, error) {
	_ = godotenv.Load(".env")
	var cfg ServerConfig
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
