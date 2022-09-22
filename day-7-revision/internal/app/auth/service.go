package auth

import (
	"day-7-revision/internal/factory"
	middlewares "day-7-revision/internal/middleware"
	"day-7-revision/internal/model"
	"day-7-revision/internal/repository"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Login(userData model.User) (string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s service) Login(userData model.User) (string, error) {
	var result model.User
	result, err := s.UserRepository.GetUserByEmailAndPassword(userData)
	if err != nil {
		return "", err
	}

	token, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
