package controller

import (
	"github.com/Nekitosss/authorization/db"
	"github.com/Nekitosss/authorization/network/controller/structures"
	"github.com/Nekitosss/authorization/utils"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"html/template"
	"bytes"
)


type EmailConfiguration interface {

	GetGmailLogin() string

	GetGmailPassword() string

	GetDomain() string

	GetGmailFrom() string
}


func AuthVerifiedUser(database *gorm.DB, info structures.AuthInfo) (*db.Session, error) {
	registerID, err := uuid.FromString(info.RegistrationID)

	if err != nil {
		return nil, err
	}

	existedUser, err := db.SelectUserByRegisterID(database, registerID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.BLErrorVerificationNotFoundUser
		} else {
			return nil, err
		}
	}

	userRegistration, err := db.SelectUserRegistration(database, registerID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.BLErrorVerificationInvalidRegID
		} else {
			return nil, err
		}
	}

	if userRegistration.Confirmed == false {
		return nil, utils.BLErrorVerificationNotConfirmedRegID
	}

	session, err := db.GetSession(database, existedUser.ID)
	return &session,  err
}


func LogIn(database *gorm.DB, info structures.LoginInfo, emailConfig EmailConfiguration) (uuid.UUID, error) {
	user, err := db.SelectUserByEmail(database, info.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		return uuid.UUID{}, err
	} else if err == gorm.ErrRecordNotFound {
		return registerNotExistedUser(database, info, emailConfig)
	} else {
		return createUserRegistration(database, user.ID, info, emailConfig)
	}
}


func registerNotExistedUser(database *gorm.DB, info structures.LoginInfo, emailConfig EmailConfiguration) (uuid.UUID, error) {
	newID := uuid.NewV4()
	newUser := db.UserModel{newID, info.Email}
	err := database.Create(&newUser).Error

	if err != nil {
		return uuid.UUID{}, err
	}

	return createUserRegistration(database, newUser.ID, info, emailConfig)
}


func createUserRegistration(database *gorm.DB, userID uuid.UUID, info structures.LoginInfo, emailConfig EmailConfiguration) (uuid.UUID, error) {
	newID := uuid.NewV4()
	newUserRegistration := db.UserRegistration{userID, newID, false}
	err := database.Create(&newUserRegistration).Error

	if err == nil {
		go sendVerificationEmail(info.Email, newUserRegistration.RegistrationID, emailConfig)
	}

	return newUserRegistration.RegistrationID, err
}


var authorizationTemplate, _ = template.ParseFiles("/root/auth_confirmation.html")


func sendVerificationEmail(toEmail string, registrationID uuid.UUID, emailConfig EmailConfiguration) {

	var link = "http://" + emailConfig.GetDomain() + "/v1/verify_register/" + registrationID.String()
	var tpl bytes.Buffer

	data := struct { RegistrationLink string } {
		RegistrationLink: link,
	}

	if err := authorizationTemplate.Execute(&tpl, data); err != nil {
		return
	}

	var signUpHTML = tpl.String()

	utils.SendEmail(emailConfig.GetGmailFrom(), emailConfig.GetGmailLogin(), emailConfig.GetGmailPassword(), []string{toEmail}, "Verification", signUpHTML)
}

