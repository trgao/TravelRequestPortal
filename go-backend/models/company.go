package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;"`
}
