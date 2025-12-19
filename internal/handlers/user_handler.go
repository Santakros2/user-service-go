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

	// This gets the params from user request
	q := r.URL.Query()

	email := q.Get("email")

	if email != "" {
		h.getUserByEmail(w, r, email)
		return
	}

	if err != nil {
		log.Println("Could not get users: ", err)
		http.Error(w, "Could not get Users", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	// For getting value from the param PathValue is used in go:1.22+
	id := r.PathValue("id")

	user, err := h.Service.GetUserById(r.Context(), id)
	if err != nil {
		http.Error(w, "Could not get user by this id", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var params models.UpdateUserDetails
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.Service.UpdateUserById(r.Context(), id, params)
	if err != nil {
		http.Error(w, "Could not Update the user ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("User was updated.")
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	userid := r.PathValue("id")

	if err := h.Service.DeleteUser(r.Context(), userid); err != nil {
		http.Error(w, "User is not present by this id.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("User deleted successfully")
}

func (h *UserHandler) getUserByEmail(w http.ResponseWriter, r *http.Request, email string) {

	user, err := h.Service.GetUserByEmail(r.Context(), email)

	if err != nil {
		http.Error(w, "No User Found with this email", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}
