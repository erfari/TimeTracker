package main

import (
	"TimeTracker/config"
	app "TimeTracker/internal/app"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("Starting TimeTracker App")
	log.Println("Initializing configuration")

	config, err := config.InitConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing database")
	dbHandler, err := app.InitDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing HTTP server")
	httpServer, err := app.InitHttpServer(config, dbHandler)
	if err != nil {
		log.Fatal(err)
	}
	httpServer.Start()
}
