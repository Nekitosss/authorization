package db

import (
	"github.com/jinzhu/gorm"

	"github.com/satori/go.uuid"
)


func (u UserModel) UpdateRegistrationID(database *gorm.DB, newRegistrationID *uuid.UUID) error {
	return database.Model(&u).Update("registration_id", newRegistrationID).Error
}


func SelectUserByLogin(database *gorm.DB, login string) (UserModel, error) {

	var user = UserModel{}
	err := database.Where("login = $1", login).First(&user).Error

	return user, err
}


func SelectUserByRegisterID(database *gorm.DB, registerID uuid.UUID) (UserModel, error) {
	var user = UserModel{}
	err := database.Where("registration_id = $1", registerID).First(&user).Error

	return user, err
}



func GetUserPasswordHash(database *gorm.DB, login string) ([]byte, error) {
	var model, err = SelectUserByLogin(database, login)

	return model.PasswordHash, err
}

func IsLoginExists(database *gorm.DB, login string) (bool, error) {
	var exists bool
	var err = database.DB().QueryRow("SELECT EXISTS(SELECT 1 FROM auth.user_model_models WHERE email = $1)", login).Scan(&exists)

	return exists, err
}
