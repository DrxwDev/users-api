package config

import (
	"os"
	"time"
)

func LoadAppConfig() AppConfig {
	return AppConfig{
		HOST:   os.Getenv("HOST"),
		PORT:   os.Getenv("PORT"),
		ACCESS: os.Getenv("ACCESS"),
	}
}

func LoadDatabase() DatabaseConfig {
	return DatabaseConfig{
		DSN: os.Getenv("DB_URL"),
	}
}

func LoadServerConfig() ServerConfig {
	return ServerConfig{
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
