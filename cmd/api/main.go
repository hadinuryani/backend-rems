package main

import (
	"log"

	"github.com/backend-rems/internal/config"
	"github.com/backend-rems/internal/database"
	"github.com/backend-rems/internal/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)

	r := router.SetupRouter(db)
	log.Println("Server menyala di port ",cfg.AppPort)
	r.Run(":"+ cfg.AppPort)
}