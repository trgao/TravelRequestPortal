package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-backend/config"
	"go-backend/models"
)

var (
	MasterConn *gorm.DB
)

func InitDatabase() {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	var err error
	MasterConn, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configure connection pool
	sqlDB, err := MasterConn.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := MasterConn.AutoMigrate(&models.Company{}, &models.User{}, &models.Request{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
