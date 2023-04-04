package application

import (
	"dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/infrastructure"
	"dasalgadoc.com/rest-websockets/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func BindRoutes(s domain.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/ping", infrastructure.PingHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/signup", infrastructure.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", infrastructure.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", infrastructure.MeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/post", infrastructure.PostSaverHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/post/{id}", infrastructure.PostFinderHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/post/{id}", infrastructure.PostUpdaterHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/post/{id}", infrastructure.PostDeleterHandler(s)).Methods(http.MethodDelete)

	r.HandleFunc("/ws", s.GetHub().RegistryClient)
}
