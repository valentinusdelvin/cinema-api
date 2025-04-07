package repository

import (
	"Cinema_API/entity"
	"Cinema_API/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUser(param models.UserLogin) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user entity.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUser(param models.UserLogin) (entity.User, error) {
	user := entity.User{}
	err := u.db.Where("email = ?", param.Email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
