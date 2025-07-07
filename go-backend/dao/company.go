package dao

import (
	"go-backend/models"

	"gorm.io/gorm"
)

func FindCompanyByName(name string, tx *gorm.DB) (company models.Company, err error) {
	result := tx.Where("name = ?", name).First(&company)
	return company, result.Error
}

func CreateCompany(name string, tx *gorm.DB) (company models.Company, err error) {
	company = models.Company{
		Name: name,
	}

	result := tx.Create(&company)
	return company, result.Error
}
