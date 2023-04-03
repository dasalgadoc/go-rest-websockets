package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"dasalgadoc.com/rest-websockets/domain"
	"encoding/json"
	"net/http"
)

func MeHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetJWTAuthorizationClaims(s, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userGetter := application.NewUserGetter(domain.UserRepositoryImplementation)
		user, err := userGetter.Do(r.Context(), claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpDto{
			Id:       string(user.Id),
			Email:    string(user.Email),
			Password: string(user.Password),
		})
	}
}
