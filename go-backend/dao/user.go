package dao

import (
	"time"

	"gorm.io/gorm"

	"go-backend/db"
	"go-backend/models"
)

func FindUserById(id uint) (user models.User, err error) {
	result := db.MasterConn.First(&user, id)
	return user, result.Error
}

func FindUserByEmail(email string) (user models.User, err error) {
	result := db.MasterConn.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetLastLoginAt(userId uint) (time time.Time, err error) {
	user := &models.User{}
	result := db.MasterConn.First(&user, userId)
	return user.LastLoginAt, result.Error
}

func UpdateLastLoginAt(userId uint) (err error) {
	user := &models.User{}
	user.ID = userId
	result := db.MasterConn.Model(&user).Update("LastLoginAt", time.Now())
	return result.Error
}

func CreateEmployee(email string, firstName string, lastName string, passwordHash string, isAdmin bool, companyId uint) (user models.User, err error) {
	user = models.User{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: passwordHash,
		IsAdmin:      isAdmin,
		LastLoginAt:  time.Now(),
		CompanyID:    companyId,
	}

	result := db.MasterConn.Create(&user)
	return user, result.Error
}

func CreateAdmin(email string, firstName string, lastName string, passwordHash string, companyId uint, tx *gorm.DB) (user models.User, err error) {
	user = models.User{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: passwordHash,
		IsAdmin:      true,
		LastLoginAt:  time.Now(),
		CompanyID:    companyId,
	}

	result := tx.Create(&user)

	return user, result.Error
}
