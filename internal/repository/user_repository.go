package repository

import (
	"context"
	"database/sql"
	"users-service/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, u models.User, passwordHash string) error {
	query := `
	INSERT INTO users (user_id, email, password_hash, name, role, active)
	VALUES (:1, :2, :3, :4, :5, :6)
	`
	_, err := r.DB.ExecContext(ctx, query,
		u.UserID,
		u.Email,
		passwordHash,
		u.Name,
		u.Role,
		u.Active,
	)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `
	SELECT user_id, email, name, role, active, created_at, updated_at
	FROM users WHERE user_id = :1
	`

	row := r.DB.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(&u.UserID, &u.Email, &u.Name, &u.Role, &u.Active, &u.CreatedAt, &u.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}
