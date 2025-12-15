package services

import (
	"context"
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
