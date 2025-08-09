package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	APIServer
}

type APIServer struct {
	Port string `env:"API_PORT" envDefault:"8080"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("error reading env: %w", err)
	}

	return &cfg, nil
}
