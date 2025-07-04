package db

import (
    "log"
    "fmt"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "go-backend/config"
    "go-backend/models"
)

var (
	MasterConn *gorm.DB
)

func InitDatabase() {
    dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", config.DBUser, config.DBPassword, config.DBPort, config.DBName)

	var err error
    MasterConn, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Println(err)
    }

    MasterConn.AutoMigrate(&models.Company{})
	MasterConn.AutoMigrate(&models.User{})
	MasterConn.AutoMigrate(&models.Request{})
}
