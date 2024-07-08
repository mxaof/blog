package login

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
