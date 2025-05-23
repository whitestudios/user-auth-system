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

func GenerateAccessJwt(email string) (string, error) {
	// Generate Access token, in this case, 15 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  "access",
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshJwt(email string) (string, error) {
	// Generating refresh token, a token with 1 week validity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  "refresh",
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
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

func ParseTokenClaims(tokenStr string) (map[string]string, error) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenStr, jwt.MapClaims{})

	if err != nil {
		return nil, err
	}

	claimsMap := make(map[string]string)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		for k, v := range claims {
			claimsMap[k] = fmt.Sprintf("%v", v)
		}
		return claimsMap, nil
	}

	return nil, fmt.Errorf("invalid claims")
}

func GenerateAccessTokenByRefreshToken(RefreshToken string) (string, error) {
	if err := VerifyToken(RefreshToken); err != nil {
		return "", fmt.Errorf("Invalid token: %v", err.Error())
	}

	claims, err := ParseTokenClaims(RefreshToken)

	if err != nil {
		return "", fmt.Errorf("invalid claims: %v", err.Error())
	}

	if claims["type"] != "refresh" {
		return "", fmt.Errorf("Invalid token")
	}

	email := claims["email"]

	token, err := GenerateAccessJwt(email)

	if err != nil {
		return "", fmt.Errorf("Error generating token")
	}

	return token, nil
}
