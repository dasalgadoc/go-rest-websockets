package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/application"
	"dasalgadoc.com/rest-websockets/domain"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"net/http"
)

type SignUpDto struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpDto{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		userCreator := application.NewUserCreator(domain.UserRepositoryImplementation)
		err = userCreator.Create(r.Context(), id.String(), request.Email, request.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		request.Id = id.String()
		request.Password = "<HIDDEN>"

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(request)
	}
}
