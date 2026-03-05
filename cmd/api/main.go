package main

import (
	"log"

	"github.com/hadinuryani/backend-rems/internal/config"
	"github.com/hadinuryani/backend-rems/internal/infrastructure/database"
	"github.com/hadinuryani/backend-rems/internal/infrastructure/logger"
	"github.com/hadinuryani/backend-rems/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.LoadConfig()

	logg := logger.NewLogger(cfg.AppEnv)
	defer logg.Sync()

	db := database.NewMySQL(cfg.DBDSN)

	_= db

	r := routes.SetupRouter(logg)
	log.Println("Server running on port",cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}