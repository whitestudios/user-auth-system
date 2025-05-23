package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// I'm using a mocked secret key (in cmd/server/main.go)
// TODO: implement .env system to get the real secret key
var secretKey []byte

func InitializeJwtService(key string) {
	secretKey = []byte(key)
}

func GenerateJwt(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
