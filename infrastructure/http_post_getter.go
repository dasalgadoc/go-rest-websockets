package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"dasalgadoc.com/rest-websockets/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type PostGetDto struct {
	Id          string
	PostContent string
	CreateAt    string
	UserId      string
}

func PostFinderHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		postGetter := application.NewPostGetter(domain.PostRepositoryImplementation)
		post, err := postGetter.Invoke(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		createAt := time.Time(post.CreatedAt)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostGetDto{
			Id:          string(post.Id),
			PostContent: string(post.PostContent),
			CreateAt:    createAt.Format(domain.DateFormat),
			UserId:      string(post.UserId),
		})
	}
}
