package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthoriseAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, ok := c.Get("IsAdmin")
		if !ok || !isAdmin.(bool) {
			log.Println("Unable to get is admin")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User is not admin",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
