package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr   string `json:"server_addr"`
	DBConnString string `json:"db_conn_string"`
	DatabaseName string `json:"database_name"`
}

func Load() (*Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	return &Config{
		ServerAddr:   os.Getenv("SERVER_ADDR"),
		DBConnString: os.Getenv("DB_CONN_STRING"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}, nil
}
