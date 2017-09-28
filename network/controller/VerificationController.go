package controller

import (
	user2 "authorization/db"
	"database/sql"

	"authorization/utils"

	"github.com/satori/go.uuid"
)

func VerifyRegistration(database *sql.DB, registerID uuid.UUID) error {
	
	var user, err = user2.SelectUserByRegisterID(database, registerID)
	
	if err != nil {

		if err == sql.ErrNoRows {
			return utils.BLErrorVerificationNotFoundUser
		} else {
			return err
		}


	}
	
	if registerID.String() != user.RegistrationID {
		return utils.BLErrorVerificationInvalidRegID
	}
	
	user.RegistrationID = ""
	
	return user.UpdateRegistrationID(database, "")
}