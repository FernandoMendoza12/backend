package service

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"auth-service/internals/domain"
	"auth-service/internals/repository"
	"auth-service/internals/utils"
)

type AuthService struct {
	repo        repository.UserRepository
	jwtSecret   string
	tokenExpiry time.Duration
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo:        repo,
		jwtSecret:   "mysecretkey",
		tokenExpiry: time.Hour * 24,
	}
}

func (s *AuthService) RegisterUser(req domain.RegisterRequest) error {

	_, err := s.repo.FindByUser(req.Username)
	if err != nil {
		return errors.New("username already exist")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	}

	return s.repo.Create(user)

}

func (s *AuthService) LoginUser(req domain.LoginRequest) (string, error) {
	user, err := s.repo.FindByUserOrEmail(req.UsernameOrEmail)
	if err != nil {
		return "", fmt.Errorf("usuario no encontrado", err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", fmt.Errorf("error generating jwtoken", err)
	}

	return token, nil
}
