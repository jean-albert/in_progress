package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString, secret string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["user_id"].(float64)), nil
	}

	return 0, jwt.ErrSignatureInvalid
}
