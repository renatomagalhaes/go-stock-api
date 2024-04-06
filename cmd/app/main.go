package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		response := healthCheckResponse{Status: "ok"}
		json.NewEncoder(w).Encode(response)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
