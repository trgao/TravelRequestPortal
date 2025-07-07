package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	ID            uint      `json:"id"`
	StartLocation string    `json:"start_location" gorm:"not null;" validate:"required"`
	Destination   string    `json:"destination" gorm:"not null;" validate:"required"`
	DateTime      time.Time `json:"date_time" gorm:"not null;" validate:"required"`
	Purpose       string    `json:"purpose" gorm:"not null;" validate:"required"`
	Status        string    `json:"status" gorm:"not null;" validate:"required"`
	AdminRemarks  string    `json:"admin_remarks" gorm:"not null;" validate:"required"`
	UserID        uint      `json:"user_id"`
	User          User      `json:"-" gorm:"not null;"`
}
