package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

var AccessSecret = []byte(os.Getenv("ACCESS_SECRET"))
var RefreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}


func GenerateAccessToken(id int, user string) (string, error) {
	return generate(id, user, 15*time.Minute, AccessSecret)
}

func generate(id int, user string, exp time.Duration, secret []byte) (string, error) {
	claims := Claims{
		UserID:   id,
		Username: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(secret)
}
