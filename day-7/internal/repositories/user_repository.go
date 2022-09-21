package repositories

import (
	"day-7/internal/entities"
)

type UserRepository interface {
	GetUsers() ([]entities.User, error)
	GetUserById(id uint) (entities.User, error)
	GetUserByEmailAndPassword(entities.User) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id uint, data entities.User) (entities.User, error)
	DeleteUser(id uint) error
}
