package Repositories

import (
	"Api/Data/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func (u UserRepository) GetUserByEmail(email string) (Models.User, error) {
	output := Models.User{}

	err := u.Table("Users").Where(&Models.User{Email: email}).First(&output)

	return output, err.Error
}

func (u UserRepository) CreateUser(user Models.User) error {
	err := u.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&user).Error
	})

	return err
}
