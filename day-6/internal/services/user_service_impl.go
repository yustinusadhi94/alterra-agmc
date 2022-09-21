package services

import (
	"day-6/internal/entities"
	"day-6/internal/repositories"
)

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u *UserServiceImpl) GetUsers() ([]entities.User, error) {
	return u.userRepo.GetUsers()
}

func (u *UserServiceImpl) GetUserById(id uint) (entities.User, error) {
	return u.userRepo.GetUserById(id)
}

func (u *UserServiceImpl) CreateUser(user entities.UserCreateRequest) (entities.User, error) {
	userData := entities.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return u.userRepo.CreateUser(userData)
}

func (u *UserServiceImpl) UpdateUser(id uint, data entities.UserUpdateRequest) (entities.User, error) {
	userData := entities.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	return u.userRepo.UpdateUser(id, userData)
}

func (u *UserServiceImpl) DeleteUser(id uint) error {
	return u.userRepo.DeleteUser(id)
}
