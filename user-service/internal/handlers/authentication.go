package handlers

import (
	"net/http"

	"github.com/necromancer26/go-microservices/user-service/internal/services"
	"github.com/necromancer26/go-microservices/user-service/internal/utils"
)

type AuthPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var authService = services.NewAuthService()

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var payload AuthPayload
	if err := utils.ParseJSONRequest(r, &payload); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	authenticated, err := authService.Authenticate(payload.Username, payload.Password)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	if authenticated {
		utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Authentication successful"})
	} else {
		utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Authentication failed"})
	}
}
