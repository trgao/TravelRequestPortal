package auth

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"go-backend/api/accessToken"
	"go-backend/config"
	"go-backend/dao"
)

type OnboardUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

type OnboardBody struct {
	User OnboardUser `json:"user"`
}

func Onboard(c *gin.Context) {
	var request OnboardBody

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is not of correct format",
		})
		return
	}

	adminIdValue, ok := c.Get("UserID")
	if !ok {
		log.Println("Unable to get user id")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Admin not authenticated",
		})
		return
	}

	adminId, ok := adminIdValue.(float64)
	if !ok {
		log.Println("Invalid admin ID type")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid admin id format",
		})
		return
	}

	admin, err := dao.FindUserById(uint(adminId))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to find admin user",
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

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	user, err = dao.CreateEmployee(request.User.Email, request.User.FirstName, request.User.LastName, string(passwordHash), request.User.IsAdmin, admin.CompanyID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create employee",
		})
		return
	}

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
		"is_admin": user.IsAdmin,
		"user_id":  user.ID,
	})
}
