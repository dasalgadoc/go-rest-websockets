package middleware

import (
	"dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/infrastructure"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

var (
	NO_AUTH_NEEDED = []string{
		"ping",
		"login",
		"signup",
	}
)

func shouldCheckToken(route string) bool {
	for _, endpoint := range NO_AUTH_NEEDED {
		if strings.Contains(route, endpoint) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s domain.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimSpace(r.Header.Get("x-auth-token"))

			_, err := jwt.ParseWithClaims(tokenString, &infrastructure.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.GetConfig().JWTSecret), nil
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
