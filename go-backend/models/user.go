package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string    `json:"email" gorm:"not null;unique;"`
	FirstName    string    `json:"first_name" gorm:"not null;"`
	LastName     string    `json:"last_name" gorm:"not null;"`
	PasswordHash string    `json:"-" gorm:"not null;"`
	IsAdmin      bool      `json:"is_admin" gorm:"not null;"`
	LastLoginAt  time.Time `json:"-" gorm:"not null;"`
	CompanyID    uint      `json:"company_id"`
	Company      Company   `json:"-" gorm:"not null;"`
}
