package main

import (
	"TimeTracker/config"
	"TimeTracker/server"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("Starting TimeTracker App")
	log.Println("Initializing configuration")

	config := config.InitConfig(".env")

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
