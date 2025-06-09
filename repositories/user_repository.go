package repositories

import (
	"go_restful_mvc/config"
	"go_restful_mvc/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmailAndPassword(email, password string) (*models.User, error)
	Update(id string, user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmailAndPassword(email, password string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(id string, user *models.User) error {
	return config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}
