package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/DrxwDev/users-api/internal/auth"
	"github.com/DrxwDev/users-api/internal/config"
	"github.com/DrxwDev/users-api/internal/database"
	"github.com/DrxwDev/users-api/internal/logger"
	"github.com/DrxwDev/users-api/internal/server"
)

func init() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found")
		}
	}
}

func main() {
	fx.New(
		config.Module,
		logger.Module,
		server.Module,
		database.Module,
		auth.Module,
	).Run()
}
