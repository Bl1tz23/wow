package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddr string `env:"SERVER_PORT" envDefault:"127.0.0.1:8080"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	Difficulty uint64 `env:"DIFFICULTY" envDefault:"1000000"`
}

func ReadConfig() (*Config, error) {
	config := Config{}

	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}

	return &config, nil
}
