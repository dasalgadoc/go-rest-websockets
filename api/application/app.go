package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/api/infrastructure"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Application struct {
	Config domain.Config
	Broker infrastructure.Broker
}

func BuildApplication() *Application {
	appConfig := getConfiguration()

	return &Application{
		Config: appConfig,
		Broker: buildBroker(appConfig),
	}
}

func getConfiguration() domain.Config {
	// load environment variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error loaging .env file", err)
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	databaseUrl := os.Getenv("DATABASE_URL")

	return domain.Config{
		Port:      port,
		JWTSecret: jwtSecret,
		Database:  databaseUrl,
	}
}

func buildBroker(config domain.Config) infrastructure.Broker {
	s, err := infrastructure.NewBroker(context.Background(), &config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return *s
}
