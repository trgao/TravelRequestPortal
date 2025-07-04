package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUser string
	DBPassword string
	DBPort string
	DBName string
	JWTSecret []byte
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}
