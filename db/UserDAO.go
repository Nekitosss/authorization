package db

import (
	"github.com/jinzhu/gorm"

	"github.com/satori/go.uuid"
)


func (u UserModel) UpdateRegistrationID(database *gorm.DB, newRegistrationID *uuid.UUID) error {
	return database.Model(&u).Update("registration_id", newRegistrationID).Error
}


func SelectUserByEmail(database *gorm.DB, email string) (UserModel, error) {
	var user = UserModel{}
	err := database.Where("email = $1", email).First(&user).Error

	return user, err
}


func SelectUserByRegisterID(database *gorm.DB, registerID uuid.UUID) (UserModel, error) {
	var user = UserModel{}
	err := database.Where("registration_id = $1", registerID).First(&user).Error

	return user, err
}



func IsUserExists(database *gorm.DB, email string) (bool, error) {
	var exists bool
	var err = database.DB().QueryRow(existsLoginValidationSQL, email).Scan(&exists)

	return exists, err
}
