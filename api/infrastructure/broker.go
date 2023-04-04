package infrastructure

import (
	"context"
	"dasalgadoc.com/rest-websockets/api/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Broker struct {
	Config *domain.Config
	Router *mux.Router
	hub    *domain.WebsocketHub
}

func NewBroker(ctx context.Context,
	config *domain.Config,
	hub *domain.WebsocketHub) (*Broker, error) {
	if err := config.ConfigErrors(); err != nil {
		return nil, err
	}

	broker := Broker{
		Config: config,
		hub:    hub,
	}

	return &broker, nil
}

func (b Broker) GetConfig() *domain.Config {
	return b.Config
}

func (b *Broker) GetHub() *domain.WebsocketHub {
	return b.hub
}

func (b *Broker) Start(binder func(s domain.Server, r *mux.Router)) {
	b.Router = mux.NewRouter()

	binder(b, b.Router)

	log.Println("Starting api on port", b.Config.Port)

	if err := http.ListenAndServe(b.Config.Port, b.Router); err != nil {
		log.Fatal("ListenAndServer ", err)
	}
}
