package infrastructure

import (
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"encoding/json"
	"net/http"
)

func PingHandler(s appDomain.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseDto{
			Message: "Pong",
			Status:  true,
		})
	}
}
