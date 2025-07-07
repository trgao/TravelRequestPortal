package accessToken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"go-backend/config"
	"go-backend/models"
)

func GenerateToken(user models.User) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       user.ID,
			"email":    user.Email,
			"is_admin": user.IsAdmin,
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		})
	return t.SignedString(config.JWTSecret)
}

func VerifyToken(tokenString string) (token *jwt.Token, err error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecret, nil
	})
}
