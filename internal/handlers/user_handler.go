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

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateUser(w, r)

	case http.MethodGet:
		h.GetAllUsers(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input models.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// calling the servie package to handle the request. Parameter is userInput model.
	user, err := h.Service.CreateUser(r.Context(), input)
	if err != nil {
		log.Println("CreateUser error:", err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers(r.Context())
	if err != nil {
		log.Println("Could not get users: ", err)
		http.Error(w, "Could not get Users", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(users)
}
