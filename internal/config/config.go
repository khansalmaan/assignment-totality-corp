package config

import (
	"os"
)

type Config struct {
	Server struct {
		Address string
	}
}

func Load() *Config {
	cfg := &Config{}
	cfg.Server.Address = getEnv("SERVER_ADDRESS", ":50051")
	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
