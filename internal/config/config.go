package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB         string
	ServerPort string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		DB:         os.Getenv("DB_CONNECTION_STRING"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
