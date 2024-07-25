package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/necromancer26/go-microservices/user-service/internal/models"
	"github.com/necromancer26/go-microservices/user-service/internal/services"
)

type AuthyPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler() *UserHandler {
	s := services.NewUserService()
	return &UserHandler{
		userService: s,
	}
}

func (u *UserHandler) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloooo")

	values, err := u.userService.GetAllUsers()

	if err != nil {
		log.Println("error getting values", err.Error())
	}
	jsonValues, err := json.Marshal(values)
	if err != nil {
		log.Println("error json values", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonValues))

}

func (u *UserHandler) UserPostHandler(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser("khalil", "chettaoui", "khalil.chettaoui@gmail")
	u.userService.CreateUser(user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
