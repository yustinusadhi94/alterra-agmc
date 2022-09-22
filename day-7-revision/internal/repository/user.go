package repository

import (
	"day-7-revision/internal/model"
	"gorm.io/gorm"
)

type User interface {
	GetUsers() ([]model.User, error)
	GetUserById(id uint) (model.User, error)
	GetUserByEmailAndPassword(model.User) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(id uint, data model.User) (model.User, error)
	DeleteUser(id uint) error
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) GetUsers() ([]model.User, error) {
	var users []model.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *user) GetUserById(id uint) (model.User, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *user) GetUserByEmailAndPassword(user model.User) (model.User, error) {
	var result model.User

	if err := r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (r *user) CreateUser(data model.User) (model.User, error) {
	tx := r.db.Begin()

	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	tx.Commit()
	return data, nil
}

func (r *user) UpdateUser(id uint, data model.User) (model.User, error) {
	_, err := r.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}

	data.ID = id

	tx := r.db.Begin()

	err = tx.Save(&data).Error
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	tx.Commit()
	return data, nil
}

func (r *user) DeleteUser(id uint) error {
	_, err := r.GetUserById(id)
	if err != nil {
		return err
	}

	tx := r.db.Begin()

	err = tx.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
