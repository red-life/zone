package cmd

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type ManagementConfig struct {
	Listen   string `env:"LISTEN,required"`
	Username string `env:"USERNAME,required"`
	Password string `env:"PASSWORD,required"`
	Postgres struct {
		Host     string `env:"HOST,required"`
		Username string `env:"USERNAME,required"`
		Password string `env:"PASSWORD,required"`
		DB       string `env:"DB,required"`
		Port     int    `env:"PORT,required"`
	} `envPrefix:"POSTGRES_"`
}

func ParseManagementConfig() ManagementConfig {
	var cfg ManagementConfig
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %s\n", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse: %s\n", err)
	}
	return cfg
}
