package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"mxaof_blog/model/login"
	"time"
)

var (
	TokenExpireDuration = time.Hour * 24
	CustomSecret        = []byte("use practice")
)

func GenToken(userId string) (string, error) {
	claims := &login.JwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "mxaof",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}
func ParseToken(tokenString string) (*login.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &login.JwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*login.JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
