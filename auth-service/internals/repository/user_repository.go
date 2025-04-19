package repository

import "auth-service/internals/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByUser(username string) (*domain.User, error)
	FindByUserOrEmail(usernameOrEmail string) (*domain.User, error)
}
