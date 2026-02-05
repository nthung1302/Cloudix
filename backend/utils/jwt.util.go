package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessSecret = []byte(os.Getenv("ACCESS_SECRET"))

type Claims struct {
	UID      string `json:"uid"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(UID string, UserName string) (string, error) {
	return generate(UID, UserName, 15*time.Minute, AccessSecret)
}

func generate(UID string, UserName string, exp time.Duration, secret []byte) (string, error) {
	claims := Claims{
		UID:      UID,
		UserName: UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(secret)
}
