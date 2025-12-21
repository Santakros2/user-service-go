package services

import (
	"context"
	"log"
	"users-service/internal/errors"
	"users-service/internal/models"
	"users-service/internal/repository"

	"github.com/google/uuid"
)

type UserService struct {
	Repo repository.DBrepo
}

func NewUserService(repo repository.DBrepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {

	id := uuid.New().String()

	// passwordHash, err := utils.HashPassword(input.Password)
	// if err != nil {
	// 	return nil, err
	// }

	log.Println(input.Email, " 	", input.Name)

	if input.Email == "" {
		return nil, errors.New(errors.CodeValidation, "email is required")
	}

	if input.Name == "" {
		return nil, errors.New(errors.CodeValidation, "name is required")
	}

	user := models.User{
		UserID: id,
		Email:  input.Email,
		Name:   input.Name,
		Role:   "USER",
		Status: "ACTIVE",
	}

	err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := s.Repo.GetAll(ctx)

	return users, err
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*models.User, error) {

	if id == "" {
		return nil, errors.New(errors.CodeValidation, "id is required")
	}
	user, err := s.Repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUserById(ctx context.Context, id string, params models.UpdateUserDetails) error {

	if id == "" {
		return errors.New(errors.CodeValidation, "id is required")
	}

	// check if there is even one value that needs to be change
	if params.Name == nil && params.Email == nil {
		return errors.New(
			errors.CodeValidation,
			"at least one field must be updated",
		)
	}

	err := s.Repo.UpdateUser(ctx, id, params)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {

	if id == "" {
		return errors.New(errors.CodeValidation, "id is required")
	}
	if err := s.Repo.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	if email == "" {
		return nil, errors.New(errors.CodeValidation, "email is required")
	}
	user, err := s.Repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
