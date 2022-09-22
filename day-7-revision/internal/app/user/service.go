package user

import (
	"day-7-revision/internal/factory"
	"day-7-revision/internal/model"
	"day-7-revision/internal/repository"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	GetUsers() ([]model.User, error)
	GetUserById(id uint) (model.User, error)
	CreateUser(user model.UserCreateRequest) (model.User, error)
	UpdateUser(id uint, data model.UserUpdateRequest) (model.User, error)
	DeleteUser(id uint) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s service) GetUsers() ([]model.User, error) {
	return s.UserRepository.GetUsers()
}

func (s service) GetUserById(id uint) (model.User, error) {
	return s.UserRepository.GetUserById(id)
}

func (s service) CreateUser(user model.UserCreateRequest) (model.User, error) {
	userData := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return s.UserRepository.CreateUser(userData)
}

func (s service) UpdateUser(id uint, data model.UserUpdateRequest) (model.User, error) {
	userData := model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	return s.UserRepository.UpdateUser(id, userData)
}

func (s service) DeleteUser(id uint) error {
	return s.UserRepository.DeleteUser(id)
}
