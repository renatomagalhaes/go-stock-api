package main

import (
	"github.com/joho/godotenv"
	"github.com/renatomagalhaes/go-stock-api/cmd/handlers"
	"github.com/renatomagalhaes/go-stock-api/database"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	defer db.Close()

	http.HandleFunc("/healthcheck", handlers.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
