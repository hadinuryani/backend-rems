package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


func NewMySQL(dsn string) *sql.DB {
	db,err := sql.Open("mysql",dsn)
	if err != nil {
		log.Fatal("Failed to open DB:",err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect DB:",err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return db
}