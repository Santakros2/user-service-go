package services

import (
	"context"
	"log"
	"users-service/internal/models"
	"users-service/internal/repository"
	"users-service/pkg/utils"

	"github.com/google/uuid"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {

	id := uuid.New().String()

	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		UserID: id,
		Email:  input.Email,
		Name:   input.Name,
		Role:   "USER",
		Active: 1,
	}

	err = s.Repo.Create(ctx, user, passwordHash)
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
	user, err := s.Repo.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUserById(ctx context.Context, id string, params models.UpdateUserDetails) error {
	err := s.Repo.UpdateUser(ctx, id, params)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	if err := s.Repo.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.Repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, err

}
