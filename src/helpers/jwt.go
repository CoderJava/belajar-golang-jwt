package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Payload struct {
	Username string
	Email    string
	Id       uint
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	jwt.RegisteredClaims
}

var JWT_SECRET string

func GenerateJwtToken(payload Payload) (string, error) {
	if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
		log.Fatal("[ERROR] JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(JWT_SECRET)

	// 7 hari
	expirationTime := time.Now().Add(24 * 60 * 7 * time.Minute)

	claims := &Claims{
		Username: payload.Username,
		Email:    payload.Email,
		Id:       payload.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expirationTime,
			},
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJwtToken(strToken string) (*Claims, error) {
	if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
		log.Fatal("[ERROR] JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(JWT_SECRET)

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, fmt.Errorf("invalid token signature")
		}
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}
