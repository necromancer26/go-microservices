package handlers

import (
	"github.com/necromancer26/go-microservices/api-gateway/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/auth", AuthHandler).Methods("POST")
	r.HandleFunc("/health", HealthHandler).Methods("GET")

	r.Use(middleware.LoggingMiddleware)

	return r
}
