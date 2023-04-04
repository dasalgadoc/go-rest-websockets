package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func PostDeleterHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetJWTAuthorizationClaims(s, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := mux.Vars(r)

		postDeleter := application.NewPostDeleter(s.GetPostRepository())
		err = postDeleter.Invoke(r.Context(), params["id"], claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseDto{
			Message: "Post deleted successfully",
			Status:  true,
		})
	}
}
