package accessToken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"go-backend/config"
	"go-backend/models"
)

func GenerateToken(user models.User) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{ 
			"id": user.ID,
			"email": user.Email,
			"is_admin": user.IsAdmin,
			"iat": time.Now().Unix(),
		})
	return t.SignedString(config.JWTSecret)
}

func VerifyToken(tokenString string) (token *jwt.Token, err error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecret, nil
   	})
}
