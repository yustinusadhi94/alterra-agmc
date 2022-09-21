package services

import (
	"day-6/internal/entities"
	"day-6/internal/middlewares"
	"day-6/internal/repositories"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (a AuthServiceImpl) Login(user entities.User) (string, error) {
	var result entities.User
	result, err := a.userRepo.GetUserByEmailAndPassword(user)
	if err != nil {
		return "", err
	}

	token, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
