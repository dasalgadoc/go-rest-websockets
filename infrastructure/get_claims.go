package infrastructure

import (
	"dasalgadoc.com/rest-websockets/api/domain"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func GetJWTAuthorizationClaims(s domain.Server, w http.ResponseWriter, r *http.Request) (*AppClaims, error) {
	tokenString := strings.TrimSpace(r.Header.Get("x-auth-token"))

	token, err := jwt.ParseWithClaims(tokenString, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.GetConfig().JWTSecret), nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return nil, nil
	}

	claims, ok := token.Claims.(*AppClaims)

	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
