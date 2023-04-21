package repositories

import "cleanCodego/internal/domain/entities"

type UserRepository interface {
	GetByID(id int64) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id int64) error
}
