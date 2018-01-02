package db

import (
	"github.com/jinzhu/gorm"

	"github.com/satori/go.uuid"
)


func SetRegistrationIDVerified(database *gorm.DB, user UserModel, registrationID uuid.UUID) error {
	verifiedRegistration := UserRegistration{user.ID, registrationID, true}
	return database.Save(&verifiedRegistration).Error
}


func SelectUserByEmail(database *gorm.DB, email string) (UserModel, error) {
	var user = UserModel{}
	err := database.Where(&UserModel{Email: email}).First(&user).Error
	return user, err
}


func SelectUserRegistration(database *gorm.DB, registerID uuid.UUID) (UserRegistration, error) {
	var userRegistration = UserRegistration{}

	err := database.
		Where(&UserRegistration{RegistrationID: registerID}).
		First(&userRegistration).
		Error

	return userRegistration, err
}


func SelectUserByRegisterID(database *gorm.DB, registerID uuid.UUID) (UserModel, error) {
	var userRegistration = UserRegistration{}
	var user = UserModel{}

	err := database.
		Where(&UserRegistration{RegistrationID: registerID}).
		First(&userRegistration).
		Error

	if err != nil {
		return user, err
	}

	err = database.
		Where(&UserModel{ID: userRegistration.UserID}).
		First(&user).
		Error

	return user, err
}



func IsUserExists(database *gorm.DB, email string) (bool, error) {
	var exists bool
	var err = database.DB().QueryRow(existsLoginValidationSQL, email).Scan(&exists)

	return exists, err
}
