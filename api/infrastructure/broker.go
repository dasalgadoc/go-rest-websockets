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
}

func NewBroker(ctx context.Context,
	config *domain.Config) (*Broker, error) {
	if err := config.ConfigErrors(); err != nil {
		return nil, err
	}

	broker := Broker{
		Config: config,
	}

	return &broker, nil
}

func (b Broker) GetConfig() *domain.Config {
	return b.Config
}

func (b *Broker) Start(binder func(s domain.Server, r *mux.Router)) {
	b.Router = mux.NewRouter()

	binder(b, b.Router)

	log.Println("Starting api on port", b.Config.Port)

	if err := http.ListenAndServe(b.Config.Port, b.Router); err != nil {
		log.Fatal("ListenAndServer ", err)
	}
}
