package handlers

import (
	"github.com/necromancer26/go-microservices/user-service/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	userHandler := NewUserHandler()
	r := mux.NewRouter()
	r.HandleFunc("/", userHandler.UserGetHandler).Methods("GET")
	r.HandleFunc("/create", userHandler.UserPostHandler).Methods("POST")
	r.HandleFunc("/auth", AuthHandler).Methods("POST")
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/service1", HealthHandler).Methods("GET")
	r.HandleFunc("/service2", HealthHandler).Methods("GET")

	r.Use(middleware.LoggingMiddleware)

	return r
}
