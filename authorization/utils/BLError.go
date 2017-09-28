package utils

import "errors"

var (

	BLErrorLoginIncorrectPassword = errors.New("BLErrorLoginIncorrectPassword")

	BLErrorUserAlreadyRegistered = errors.New("BLErrorUserAlreadyRegistered")

	BLErrorVerificationNotFoundUser = errors.New("BLErrorVerificationNotFoundUser")

	BLErrorVerificationInvalidRegID = errors.New("BLErrorVerificationInvalidRegID")

)
