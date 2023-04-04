package infrastructure

import (
	"context"
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Broker struct {
	Config         *appDomain.Config
	Router         *mux.Router
	hub            *appDomain.WebsocketHub
	userRepository domain.UserRepository
	postRepository domain.PostRepository
}

func NewBroker(ctx context.Context,
	config *appDomain.Config,
	hub *appDomain.WebsocketHub,
	u domain.UserRepository,
	p domain.PostRepository) (*Broker, error) {
	if err := config.ConfigErrors(); err != nil {
		return nil, err
	}

	broker := Broker{
		Config:         config,
		hub:            hub,
		userRepository: u,
		postRepository: p,
	}

	return &broker, nil
}

func (b Broker) GetConfig() *appDomain.Config {
	return b.Config
}

func (b *Broker) GetHub() *appDomain.WebsocketHub {
	return b.hub
}

func (b *Broker) GetUserRepository() domain.UserRepository {
	return b.userRepository
}

func (b *Broker) GetPostRepository() domain.PostRepository {
	return b.postRepository
}

func (b *Broker) Start(binder func(s appDomain.Server, r *mux.Router)) {
	b.Router = mux.NewRouter()

	binder(b, b.Router)

	log.Println("Starting api on port", b.Config.Port)

	if err := http.ListenAndServe(b.Config.Port, b.Router); err != nil {
		log.Fatal("ListenAndServer ", err)
	}
}
