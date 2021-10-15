package main

import (
	"log"
	"os"

	"github.com/philvc/jobbi-api/config"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/database"
	"github.com/philvc/jobbi-api/repository"
	"github.com/philvc/jobbi-api/router"
	"github.com/philvc/jobbi-api/usecase"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	godotenv.Load()

	// Creates the config
	config := config.Config{
		Database: config.Database{
			ConnectionString: os.Getenv("DATABASE_URL"),
		},
	}

	// Config sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://eb45bee00df04e23b5bf406fec26a968@o498150.ingest.sentry.io/5996194",
		// Specify a fixed sample rate:
		TracesSampleRate: 1,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	sentry.CaptureMessage("It works!")

	// Initialises the router
	engine := gin.Default()

	// Connects to the database
	database := database.Default(config)

	// Creates the repository container
	repository := repository.Default(database)

	// Creates the usecase container
	usecase := usecase.Default(repository)

	// Creates the controller container
	controller := controller.Default(usecase)

	// Connect routes & start router
	router := router.Default(engine, controller)

	// Initialises the router
	router.Initiliase()

	// Start the router
	router.Run()
}
