package auth

import (
	"go-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", config.Host, false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
