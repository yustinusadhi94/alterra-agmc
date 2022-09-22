package factory

import (
	"day-7-revision/database"
	"day-7-revision/internal/repository"
)

type Factory struct {
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		UserRepository: repository.NewUserRepository(db),
	}
}
