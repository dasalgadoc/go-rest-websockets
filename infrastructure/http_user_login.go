package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = LoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userLogger := application.NewUserLogin(s.GetUserRepository())
		user, err := userLogger.Invoke(r.Context(), request.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.IsUserEmpty() {
			http.Error(w, "No valid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			http.Error(w, "No valid credentials", http.StatusUnauthorized)
			return
		}

		// JWT!
		claims := AppClaims{
			UserId: string(user.Id),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		stringToken, err := token.SignedString([]byte(s.GetConfig().JWTSecret))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{Token: stringToken})

	}
}
