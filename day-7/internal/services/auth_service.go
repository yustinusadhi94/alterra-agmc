package services

import "day-7/internal/entities"

type AuthService interface {
	Login(entities.User) (string, error)
}
