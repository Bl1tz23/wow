package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddr     string `env:"SERVER_ADDR" envDefault:"127.0.0.1:8080"`
	RequestsToMake int    `env:"REQUESTS_TO_MAKE" envDefault:"50"`
}

func ReadConfig() (*Config, error) {
	config := Config{}

	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}

	return &config, nil
}
