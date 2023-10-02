package common

import (
	"github.com/China-song/ginEssential/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 过期时间 一周后过期
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "github.com/China-song",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
