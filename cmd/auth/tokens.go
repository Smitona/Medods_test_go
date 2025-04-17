package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func generateAccessToken(GUID string, IPAddress string) (string, error) {
	// Generate JWT token with GUID and IP address as claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": GUID,
		"ip":   IPAddress,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateRefreshToken() (string, error) {
	// Generating Refresh Token (random bytes, base64 encoded)
	str := make([]byte, 32)
	_, err := rand.Read(str)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(str), nil
}
