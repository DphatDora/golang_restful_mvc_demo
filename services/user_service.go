package services

import (
	"go_restful_mvc/models"
	"go_restful_mvc/repositories"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	Update(id string, user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	return s.repo.FindByEmailAndPassword(email, password)
}

func (s *userService) Update(id string, user *models.User) error {
	return s.repo.Update(id, user)
}
