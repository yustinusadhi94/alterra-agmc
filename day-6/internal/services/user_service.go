package services

import "day-6/internal/entities"

type UserService interface {
	GetUsers() ([]entities.User, error)
	GetUserById(id uint) (entities.User, error)
	CreateUser(user entities.UserCreateRequest) (entities.User, error)
	UpdateUser(id uint, data entities.UserUpdateRequest) (entities.User, error)
	DeleteUser(id uint) error
}
