package main

import (
	"log"
	"net/http"

	"github.com/necromancer26/go-microservices/user-service/internal/handlers"
)

func main() {
	r := handlers.SetupRouter()
	log.Println("Starting API Gateway on port 8081...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
