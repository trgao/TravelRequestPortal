package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Host       string
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBName     string
	JWTSecret  []byte
	CORSOrigin string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	CORSOrigin = os.Getenv("CORS_ORIGIN")
}
