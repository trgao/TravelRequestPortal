package middlewares

import (
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthoriseUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			c.Abort()
			return
		}

		userIdValue, ok := c.Get("UserID")
		if !ok {
			log.Println("Unable to get user id")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User does not have permission to view requests",
			})
			c.Abort()
			return
		}

		userId, ok := userIdValue.(float64)
		if !ok || int64(userId) != queryId {
			log.Println("Invalid user id")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User does not have permission to view requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
