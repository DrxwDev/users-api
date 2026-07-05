package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/DrxwDev/users-api/internal/auth"
	"github.com/DrxwDev/users-api/internal/config"
	"github.com/DrxwDev/users-api/internal/database"
	"github.com/DrxwDev/users-api/internal/logger"
	"github.com/DrxwDev/users-api/internal/server"
	"github.com/DrxwDev/users-api/internal/users"
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
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	fx.New(
		config.Module,
		logger.Module,
		server.Module,
		database.Module,
		auth.Module,
		users.Module,
	).Run()
}
