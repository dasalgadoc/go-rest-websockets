package application

import (
	"dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/infrastructure"
	"github.com/gorilla/mux"
	"net/http"
)

func BindRoutes(s domain.Server, r *mux.Router) {
	r.HandleFunc("/ping", infrastructure.PingHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", infrastructure.SignUpHandler(s)).Methods(http.MethodPost)
}
