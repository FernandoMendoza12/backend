package auth

import (
	"time"

	"auth-service/internals/repository"

)

type AuthService struct {
	repo        repository.UserRepository
	jwt         string
	tokenExpiry time.Duration
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repository.UserRepository,
	}
}
