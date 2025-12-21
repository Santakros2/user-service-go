package models

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateUserDetails struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}
