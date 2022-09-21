package services

import "day-6/internal/entities"

type AuthService interface {
	Login(entities.User) (string, error)
}
