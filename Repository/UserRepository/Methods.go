package UserRepository

import (
	"Api/Config"
	"Api/Data/Models"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) *Models.User {
	var existing Models.User
	err := Config.DB.Preload("ClinicalUser.Clinical").Where("email = ?", email).First(&existing).Error
	if err == nil {
		return nil
	}
	return &existing
}

func RegisterUser(user Models.User, tx *gorm.DB) *error {
	db := tx
	if tx == nil {
		db = Config.DB
	}

	err := db.Create(&user).Error
	if err == nil {
		return nil
	}

	return &err
}
