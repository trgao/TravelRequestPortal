package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"

	"go-backend/api/accessToken"
	"go-backend/dao"
)

func AuthenticateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "There is no access token",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := accessToken.VerifyToken(cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid access token",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid access token",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if float64(time.Now().Unix()) > claims["iat"].(float64)+86400 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access token expired",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		lastLoginAt, err := dao.GetLastLoginAt(uint(claims["id"].(float64)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if float64(lastLoginAt.Unix()) > claims["iat"].(float64)*1000 {
			c.SetCookie("token", "", 86400, "/", "localhost", false, true)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access token expired",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("UserID", claims["id"])
		c.Set("IsAdmin", claims["is_admin"])
		c.Next()
	}
}
