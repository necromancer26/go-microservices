package middleware

import (
	"net/http"
	"strings"

	"github.com/necromancer26/go-microservices/user-service/internal/utils"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Missing Authorization header"})
			return
		}

		// Example of a simple token validation
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 || parts[1] != "the-token-to-implement-when-i-have-time" {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid or missing token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
