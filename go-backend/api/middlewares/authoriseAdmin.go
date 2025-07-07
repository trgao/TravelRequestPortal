package middlewares

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthoriseAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, ok := c.Get("IsAdmin")
		if !ok {
			log.Println("Unable to get IsAdmin")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User is not admin",
			})
			c.Abort()
			return
		}

		isAdminBool, ok := isAdmin.(bool)
		if !ok || !isAdminBool {
			log.Println("User is not admin")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User is not admin",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
