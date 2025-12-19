package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"users-service/internal/models"
)

type DBrepo interface {
	Create(ctx context.Context, u models.User, passwordHash string) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, updatedata models.UpdateUserDetails) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, u models.User, passwordHash string) error {
	query := `
	INSERT INTO users (user_id, email, password_hash, name, role, active)
	VALUES (?, ?, ?, ?, ?, ?)
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
	FROM users WHERE user_id = ?
	`

	row := r.DB.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(&u.UserID, &u.Email, &u.Name, &u.Role, &u.Active, &u.CreatedAt, &u.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0)

	query := `SELECT * FROM USERS`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var skip any

	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(
			&user.UserID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
			&skip,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	query := `select user_id, name, email, role, active from users where id = ?`

	row := r.DB.QueryRowContext(ctx, query, id)
	var user models.User
	if err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Role, &user.Active); err == sql.ErrNoRows {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, updatedata models.UpdateUserDetails) error {

	clauses := []string{}
	args := []any{}

	if updatedata.Name != nil {
		clauses = append(clauses, "name= ?")
		args = append(args, updatedata.Name)
	}

	if updatedata.Email != nil {
		clauses = append(clauses, "email=?")
		args = append(args, updatedata.Email)
	}

	if len(clauses) == 0 {
		return errors.New("no fields to update")
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE user_id = ?", strings.Join(clauses, ", "))

	args = append(args, id)

	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE user_id = ? `

	_, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT user_id, name, email, role, active FROM users WHERE email = ?`

	row := r.DB.QueryRowContext(ctx, query, email)

	var user models.User

	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Role, &user.Active)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	// log.Println(user)
	return &user, nil
}
