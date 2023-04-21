package database

import (
	"cleanCodego/internal/domain/entities"
	"cleanCodego/internal/domain/repositories"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id int64) (*entities.User, error) {
	// Implement database logic here
	return nil, nil
}

func (r *userRepository) Create(user *entities.User) error {
	// Implement database logic here
	return nil
}

func (r *userRepository) Update(user *entities.User) error {
	// Implement database logic here
	return nil
}

func (r *userRepository) Delete(id int64) error {
	// Implement database logic here
	return nil
}
