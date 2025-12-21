package handlers

import (
	"encoding/json"
	stderr "errors"
	"net/http"
	"users-service/internal/errors"
	"users-service/internal/logger"
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
		h.writeError(w, errors.New(
			errors.CodeValidation,
			"invalid request body",
		))
		return
	}

	// calling the servie package to handle the request. Parameter is userInput model.
	user, err := h.Service.CreateUser(r.Context(), input)
	if err != nil {
		h.writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	// This gets the params from user request
	q := r.URL.Query()

	email := q.Get("email")

	if email != "" {
		h.getUserByEmail(w, r, email)
		return
	}
	users, err := h.Service.GetAllUsers(r.Context())

	if err != nil {
		h.writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	// For getting value from the param PathValue is used in go:1.22+
	id := r.PathValue("id")

	user, err := h.Service.GetUserById(r.Context(), id)
	if err != nil {
		h.writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var params models.UpdateUserDetails
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		h.writeError(w, errors.New(
			errors.CodeValidation,
			"invalid request body",
		))
		return
	}

	err := h.Service.UpdateUserById(r.Context(), id, params)
	if err != nil {
		h.writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode("User was updated.")
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	userid := r.PathValue("id")

	if err := h.Service.DeleteUser(r.Context(), userid); err != nil {
		h.writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode("User deleted successfully")
}

func (h *UserHandler) getUserByEmail(w http.ResponseWriter, r *http.Request, email string) {

	user, err := h.Service.GetUserByEmail(r.Context(), email)

	if err != nil {
		h.writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) writeError(w http.ResponseWriter, err error) {
	logger.Logger.Println("request error:", err)

	var appErr *errors.AppError

	if !stderr.As(err, &appErr) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	switch appErr.Code {
	case errors.CodeValidation:
		http.Error(w, appErr.Message, http.StatusBadRequest)
	case errors.CodeConflict:
		http.Error(w, appErr.Message, http.StatusConflict)
	case errors.CodeUnauthorized:
		http.Error(w, appErr.Message, http.StatusUnauthorized)
	case errors.CodeForbidden:
		http.Error(w, appErr.Message, http.StatusForbidden)
	case errors.CodeNotFound:
		http.Error(w, appErr.Message, http.StatusNotFound)
	case errors.CodeTimeout:
		http.Error(w, appErr.Message, http.StatusGatewayTimeout)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

}
