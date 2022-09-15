package database

import (
	"day-3/config"
	"day-3/middlewares"
	"day-3/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func GetUserById(id uint) (models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(data models.User) (models.User, error) {
	tx := config.DB.Begin()

	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return data, nil
}

func UpdateUser(id uint, data models.User) (models.User, error) {
	_, err := GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	data.ID = id

	tx := config.DB.Begin()

	err = tx.Save(&data).Error
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return data, nil
}

func DeleteUser(id uint) error {
	_, err := GetUserById(id)
	if err != nil {
		return err
	}

	tx := config.DB.Begin()

	err = tx.Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func LoginUser(user models.User) (string, error) {
	var result models.User
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&result).Error
	if err != nil {
		return "", err
	}

	token, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
