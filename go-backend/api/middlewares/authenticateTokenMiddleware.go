package middlewares

import (
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-backend/api/accessToken"
	"go-backend/config"
	"go-backend/dao"
)

func AuthenticateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "There is no access token",
			})
			c.Abort()
			return
		}

		token, err := accessToken.VerifyToken(cookie)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid access token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Cannot get token claims")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid access token",
			})
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access token expired",
			})
			c.Abort()
			return
		}

		lastLoginAt, err := dao.GetLastLoginAt(uint(claims["id"].(float64)))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			c.Abort()
			return
		}

		if float64(lastLoginAt.Unix()) > claims["iat"].(float64) {
			c.SetCookie("token", "", 86400, "/", config.Host, false, true)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access token expired",
			})
			c.Abort()
			return
		}

		c.Set("UserID", claims["id"])
		c.Set("IsAdmin", claims["is_admin"])
		c.Next()
	}
}
