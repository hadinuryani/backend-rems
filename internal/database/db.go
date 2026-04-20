package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/backend-rems/internal/config"
)

func Connect(cfg *config.Config) *sql.DB{
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db,err := sql.Open("mysql",dsn)
	if err != nil{
		log.Fatal("Gagal Membuat koneksi : ",err)
	}

	if err := db.Ping();err != nil {
		log.Fatal("Gagal Terkoneksi Database",err)
	}
	// polling e
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	
	return db
}