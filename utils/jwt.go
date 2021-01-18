package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go_pro1/model"
	"time"
)

var jwtKey = []byte("a_secret_crect)")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error)  {
	expirationTime := time.Now().Add((7 * 24 * time.Hour))

	claims := Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "ken.by.gotest",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
