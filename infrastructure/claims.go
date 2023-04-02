package infrastructure

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
