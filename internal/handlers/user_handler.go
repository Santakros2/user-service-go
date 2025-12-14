package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"users-service/internal/models"
	"users-service/internal/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input models.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.Service.CreateUser(r.Context(), input)
	if err != nil {
		log.Println("CreateUser error:", err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)
}

// func (h *UserHandler) GetUsers(w http.ResponseWriter, r http.Request) {
// 	if (r.)
// }
