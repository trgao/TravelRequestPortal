package auth

import (
	"log"

	"github.com/absagar/go-bcrypt"
	"github.com/gin-gonic/gin"
	"net/http"

	"go-backend/api/accessToken"
	"go-backend/dao"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginBody struct {
	User LoginUser `json:"user"`
}

func Login(c *gin.Context) {
	var request LoginBody

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Request body is not of correct format",
		})
		return
	}

	user, err := dao.FindUserByEmail(request.User.Email)
	if err != nil || !bcrypt.Match(request.User.Password, user.PasswordHash) {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password is not correct",
		})
		return
	}

	err = dao.UpdateLastLoginAt(user.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot update last login at",
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

	c.SetCookie("token", token, 86400, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully logged in",
		"is_admin": user.IsAdmin,
		"user_id":  user.ID,
	})
}
