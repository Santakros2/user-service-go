package repository

import (
	"context"
	"database/sql"
	stderr "errors"
	"fmt"
	"log"
	"strings"
	"users-service/internal/errors"
	"users-service/internal/models"

	"github.com/go-sql-driver/mysql"
)

type DBrepo interface {
	Create(ctx context.Context, u models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
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

func (r *UserRepository) Create(ctx context.Context, u models.User) error {
	query := `
	INSERT INTO users (user_id, email, name, role, status)
	VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.DB.ExecContext(ctx, query,
		u.UserID,
		u.Email,
		u.Name,
		u.Role,
		u.Status,
	)

	// error handling
	if err != nil {
		if stderr.Is(err, context.Canceled) || stderr.Is(err, context.DeadlineExceeded) {
			return errors.Wrap(errors.CodeTimeout, "request timeout", err)
		}

		// the error DB internally can generate mulple type
		// Like Timeout, Cancelled context, IO error, network error, mysql error
		// mysql because i am using mysql otherwise specific db error that is being used
		var mysqlErr *mysql.MySQLError
		if stderr.As(err, &mysqlErr) && mysqlErr != nil {
			if mysqlErr.Number == 1062 {
				return errors.New(
					errors.CodeConflict,
					"email already exists",
				)
			}
		}

		return errors.Wrap(errors.CodeInternal, "database error", err)
	}

	rows, err := result.RowsAffected()
	if err == nil {
		log.Println("ROWS INSERTED:", rows)
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `
	SELECT user_id, email, name, role, status, created_at, updated_at
	FROM users WHERE user_id = ?
	`

	row := r.DB.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(&u.UserID, &u.Email, &u.Name, &u.Role, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if stderr.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.CodeNotFound, "user not found")
		}
		return nil, errors.Wrap(errors.CodeInternal, "database error", err)
	}
	return &u, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0)

	query := `SELECT user_id, name, email, role, status, created_at, updated_at FROM users `

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(
			errors.CodeInternal,
			"database error",
			err,
		)
	}

	defer rows.Close()

	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(
			&user.UserID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(
				errors.CodeInternal,
				"row scan error",
				err,
			)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(
			errors.CodeInternal,
			"row iteration error",
			err,
		)
	}

	return users, nil

}

// func (r *UserRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
// 	query := `select user_id, name, email, role, active from users where user_id = ?`

// 	row := r.DB.QueryRowContext(ctx, query, id)
// 	var user models.User
// 	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Role, &user.Status)
// 	if err != nil {
// 		if stderr.Is(err, sql.ErrNoRows) {
// 			return nil, errors.New(errors.CodeNotFound, "user id not found")
// 		}
// 		return nil, errors.Wrap(errors.CodeInternal, "database error", err)
// 	}

// 	return &user, nil
// }

func (r *UserRepository) UpdateUser(ctx context.Context, id string, updatedata models.UpdateUserDetails) error {

	clauses := []string{}
	args := []any{}

	if updatedata.Name != nil {
		clauses = append(clauses, "name= ?")
		args = append(args, *updatedata.Name)
	}

	if updatedata.Email != nil {
		clauses = append(clauses, "email=?")
		args = append(args, *updatedata.Email)
	}

	if len(clauses) == 0 {
		return errors.New(
			errors.CodeValidation,
			"no fields to update",
		)
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE user_id = ?", strings.Join(clauses, ", "))

	args = append(args, id)

	result, err := r.DB.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.Wrap(
			errors.CodeInternal,
			"update failed",
			err,
		)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(
			errors.CodeInternal,
			"failed to check update result",
			err,
		)
	}

	if rows == 0 {
		return errors.New(
			errors.CodeNotFound,
			"user not found",
		)
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE user_id = ? `

	result, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return errors.Wrap(
			errors.CodeInternal,
			"delete failed",
			err,
		)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(
			errors.CodeInternal,
			"failed to check update result",
			err,
		)
	}

	if rows == 0 {
		return errors.New(
			errors.CodeNotFound,
			"user not found",
		)
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT user_id, name, email, role, status FROM users WHERE email = ?`

	row := r.DB.QueryRowContext(ctx, query, email)

	var user models.User

	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Role, &user.Status)

	if err != nil {
		if stderr.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.CodeNotFound, "user not found")
		}
		return nil, errors.Wrap(errors.CodeInternal, "database error", err)
	}

	// log.Println(user)
	return &user, nil
}
