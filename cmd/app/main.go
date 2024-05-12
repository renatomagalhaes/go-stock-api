package main

import (
	"github.com/renatomagalhaes/go-stock-api/cmd/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", handlers.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
