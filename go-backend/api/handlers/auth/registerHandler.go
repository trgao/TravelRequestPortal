package auth

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"go-backend/api/accessToken"
	"go-backend/config"
	"go-backend/dao"
	"go-backend/db"
)

type RegisterUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Company   string `json:"company"`
}

type RegisterBody struct {
	User RegisterUser `json:"user"`
}

func Register(c *gin.Context) {
	var request RegisterBody

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is not of correct format",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	user, findErr := dao.FindUserByEmail(request.User.Email)
	if findErr == nil {
		log.Println(findErr)
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}

	tx := db.MasterConn.Begin()
	defer tx.Rollback()

	company, findErr := dao.FindCompanyByName(request.User.Company, tx)
	if findErr == nil {
		log.Println(findErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Company already exists",
		})
		return
	}

	company, err = dao.CreateCompany(request.User.Company, tx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create company",
		})
		return
	}

	user, err = dao.CreateAdmin(request.User.Email, request.User.FirstName, request.User.LastName, string(passwordHash), company.ID, tx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create admin",
		})
		return
	}

	tx.Commit()

	token, err := accessToken.GenerateToken(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to generate token",
		})
		return
	}

	c.SetCookie("token", token, 86400, "/", config.Host, false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully registered a new account",
		"is_admin": true,
		"user_id":  user.ID,
	})
}
