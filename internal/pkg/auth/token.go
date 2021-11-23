package auth

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

// GenerateToken generate jwt token
func GenerateToken(userId string) (string, error) {
	// Always store the access secret for jwt token in a secret
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, createClaims(userId))
	token, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func createClaims(userId string) jwt.MapClaims {
	return jwt.MapClaims{
		"authorized": true,
		"expiry":     time.Now().Add(15 * time.Minute),
		"user_id":    userId,
	}
}
