package main

import (
	"log"
	"net/http"

	"github.com/necromancer26/go-microservices/api-gateway/config"
	"github.com/necromancer26/go-microservices/api-gateway/internal/handlers"
	"github.com/necromancer26/go-microservices/api-gateway/pkg/logger"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()

	r := handlers.SetupRouter()

	log.Println("Starting API Gateway on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
