package config

import (
	"os"

	"github.com/joho/godotenv"
)

var AppConfig, MysqlConfig map[string]interface{}

func init() {
	godotenv.Load()
	AppConfig = map[string]interface{}{
		"port": os.Getenv("SERVER_PORT"),
		"host": os.Getenv("SERVER_ADDRESS"),
	}

	MysqlConfig = map[string]interface{}{
		"username": os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASS"),
		"database": os.Getenv("DB_NAME"),
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
	}
}
