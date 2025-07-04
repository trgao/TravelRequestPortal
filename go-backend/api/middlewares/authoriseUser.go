package middlewares

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthoriseUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		userId, ok := c.Get("UserID")
		if !ok || int64(userId.(float64)) != queryId {
			log.Println("Unable to get user id")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User does not have permission to view requests",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
