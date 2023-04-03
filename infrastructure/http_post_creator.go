package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"dasalgadoc.com/rest-websockets/domain"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"net/http"
	"time"
)

type PostDto struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

func PostSaver(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetJWTAuthorizationClaims(s, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var request = PostDto{}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		postCreator := application.NewPostCreator(domain.PostRepositoryImplementation)
		err = postCreator.Invoke(r.Context(), id.String(), request.PostContent, claims.UserId, time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostResponse{
			Id:          id.String(),
			PostContent: request.PostContent,
		})
	}
}
