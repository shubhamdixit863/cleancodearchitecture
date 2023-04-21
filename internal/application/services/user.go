package services

import (
	"cleanCodego/internal/domain/entities"
	"cleanCodego/internal/domain/repositories"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) GetByID(id int64) (*entities.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Create(user *entities.User) error {
	return s.repo.Create(user)
}

func (s *userService) Update(user *entities.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id int64) error {
	return s.repo.Delete(id)
}
