package middleware

import (
	"log"
	"net/http"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

func ErrorWrapper(handler HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			log.Println("Handler error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
