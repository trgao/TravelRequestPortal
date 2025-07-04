package dao

import (
	"go-backend/db"
	"go-backend/models"
)

func FindCompanyByName(name string) (company models.Company, err error) {
	result := db.MasterConn.Where("name = ?", name).First(&company)
	return company, result.Error
}

func CreateCompany(name string) (company models.Company, err error) {
	company = models.Company {
		Name: name,
	}

	result := db.MasterConn.Create(&company)
	return company, result.Error
}
