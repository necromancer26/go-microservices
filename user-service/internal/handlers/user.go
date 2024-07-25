package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/necromancer26/go-microservices/user-service/internal/models"
	"github.com/necromancer26/go-microservices/user-service/internal/services"
)

//	type AuthyPayload struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
//
//	type UserPayload struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler() *UserHandler {
	s := services.NewUserService()
	return &UserHandler{
		userService: s,
	}
}

func (u *UserHandler) UserGetHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(users)
}

func (u *UserHandler) UserPostHandler(w http.ResponseWriter, r *http.Request) error {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}

	err = u.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	return nil
}
func (u *UserHandler) UserGetByNameHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	name := vars["name"]
	user, err := u.userService.GetUserByName(name)
	if err != nil {
		http.Error(w, "Could not fetch user", http.StatusInternalServerError)
		return err
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) UserDeleteHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return err
	}

	err = u.userService.DeleteUserById(id)
	if err != nil {
		http.Error(w, "Could not delete user", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
	return nil
}
func (u *UserHandler) UserUpdateHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return err
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}

	user.ID = id
	err = u.userService.UpdateUser(&user)
	if err != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
	return nil
}
