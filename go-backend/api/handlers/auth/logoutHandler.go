package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
    c.SetCookie("token", "", 86400, "/", "localhost", false, true)
    c.JSON(http.StatusOK, gin.H{
        "message": "Successfully logged out",
    })
}
