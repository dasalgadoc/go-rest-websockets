package application

import (
	"context"
	appDomain "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/api/infrastructure"
	"dasalgadoc.com/rest-websockets/domain"
	"dasalgadoc.com/rest-websockets/infrastructure/repository"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type (
	Application struct {
		Config appDomain.Config
		Broker infrastructure.Broker
	}

	applicationRepositories struct {
		userRepository domain.UserRepository
		postRepository domain.PostRepository
	}
)

func BuildApplication() *Application {
	appConfig := getConfiguration()

	repositories := buildRepositories(appConfig)

	hub := buildWebSocket()

	return &Application{
		Config: appConfig,
		Broker: buildBroker(appConfig, hub, repositories),
	}
}

func getConfiguration() appDomain.Config {
	// load environment variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error loaging .env file", err)
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	databaseUrl := os.Getenv("DATABASE_URL")

	return appDomain.Config{
		Port:      port,
		JWTSecret: jwtSecret,
		Database:  databaseUrl,
	}
}

func buildBroker(config appDomain.Config, hub *appDomain.WebsocketHub, repositories applicationRepositories) infrastructure.Broker {
	s, err := infrastructure.NewBroker(context.Background(), &config, hub, repositories.userRepository, repositories.postRepository)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return *s
}

func buildRepositories(config appDomain.Config) applicationRepositories {
	// If we have more than one implementation on UserRepository a builder is in order (with .env)
	users, err := repository.NewPostgresUserRepository(config.Database)
	if err != nil {
		log.Fatalln(err.Error())
	}

	posts, err := repository.NewPostgresPostRepository(config.Database)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return applicationRepositories{
		userRepository: users,
		postRepository: posts,
	}
}

func buildWebSocket() *appDomain.WebsocketHub {
	hub := appDomain.NewHub()
	go hub.Run()
	return hub
}
