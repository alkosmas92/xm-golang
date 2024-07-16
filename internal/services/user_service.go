package services

import (
	"context"
	"errors"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *models.User) error
	AuthenticateUser(ctx context.Context, username, password string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, user *models.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	log.Print("user: ", user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
