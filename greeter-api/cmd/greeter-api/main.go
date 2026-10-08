package main

import (
	"log"
	"net/http"

	"greeter-api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.Greeting)

	log.Println("greeter-api listening on :9090")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatal(err)
	}
}
