package controller

import (
	user2 "github.com/Nekitosss/authorization/db"
	"github.com/jinzhu/gorm"

	"github.com/Nekitosss/authorization/utils"

	"github.com/satori/go.uuid"
)


func VerifyRegistration(database *gorm.DB, registerID uuid.UUID) error {
	
	var user, err = user2.SelectUserByRegisterID(database, registerID)
	
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return utils.BLErrorVerificationNotFoundUser
		} else {
			return err
		}
	}
	
	if registerID != *user.RegistrationID {
		return utils.BLErrorVerificationInvalidRegID
	}
	
	user.RegistrationID = nil
	
	return user.UpdateRegistrationID(database, nil)
}