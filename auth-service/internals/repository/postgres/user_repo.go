package postgres

import (
	"gorm.io/gorm"

	"auth-service/internals/domain"
	"auth-service/internals/repository"

)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {

	db.AutoMigrate(&domain.User{})
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUser(username string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("username=?", username).First(&user)
	if result != nil {
		return nil, result.Error
	}
	return &user, nil
}
