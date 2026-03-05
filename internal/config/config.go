package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string
	AppEnv  string
	DBDSN   string
}

func LoadConfig() *Config {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		AppEnv:  os.Getenv("APP_ENV"),
		DBDSN:   dsn,
	}
}
