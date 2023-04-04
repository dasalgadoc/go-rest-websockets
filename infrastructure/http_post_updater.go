package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func PostUpdaterHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetJWTAuthorizationClaims(s, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := mux.Vars(r)

		var request = PostDto{}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		postUpdater := application.NewPostUpdater(s.GetPostRepository())
		err = postUpdater.Invoke(r.Context(), params["id"], request.PostContent, claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseDto{
			Message: "Post updated successfully",
			Status:  true,
		})
	}
}
