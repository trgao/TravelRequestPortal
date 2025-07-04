package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	ID            uint      `json:"id"`
	StartLocation string    `json:"start_location" gorm:"not null;"`
	Destination   string    `json:"destination" gorm:"not null;"`
	DateTime      time.Time `json:"date_time" gorm:"not null;"`
	Purpose       string    `json:"purpose" gorm:"not null;"`
	Status        string    `json:"status" gorm:"not null;"`
	AdminRemarks  string    `json:"admin_remarks" gorm:"not null;"`
	UserID        uint      `json:"user_id"`
	User          User      `json:"-" gorm:"not null;"`
}
