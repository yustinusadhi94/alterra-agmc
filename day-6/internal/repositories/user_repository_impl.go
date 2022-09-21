package repositories

import (
	"day-6/internal/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetUsers() ([]entities.User, error) {
	var users []entities.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetUserById(id uint) (entities.User, error) {
	var user entities.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByEmailAndPassword(user entities.User) (entities.User, error) {
	var result entities.User

	if err := r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (r *UserRepositoryImpl) CreateUser(data entities.User) (entities.User, error) {
	tx := r.db.Begin()

	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return entities.User{}, err
	}

	tx.Commit()
	return data, nil
}

func (r *UserRepositoryImpl) UpdateUser(id uint, data entities.User) (entities.User, error) {
	_, err := r.GetUserById(id)
	if err != nil {
		return entities.User{}, err
	}

	data.ID = id

	tx := r.db.Begin()

	err = tx.Save(&data).Error
	if err != nil {
		tx.Rollback()
		return entities.User{}, err
	}

	tx.Commit()
	return data, nil
}

func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	_, err := r.GetUserById(id)
	if err != nil {
		return err
	}

	tx := r.db.Begin()

	err = tx.Where("id = ?", id).Delete(&entities.User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
