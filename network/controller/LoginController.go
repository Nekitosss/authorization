package controller

import (
	"github.com/Nekitosss/authorization/db"
	"github.com/Nekitosss/authorization/network/controller/structures"
	"github.com/Nekitosss/authorization/utils"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)


type EmailConfiguration interface {

	GetGmailLogin() string

	GetGmailPassword() string

	GetDomain() string
}


func LogIn(database *gorm.DB, info structures.LoginInfo, emailConfig EmailConfiguration) (*db.Session, uuid.NullUUID, error) {
	exists, err := db.IsUserExists(database, info.Email)

	if err != nil {
		return nil, uuid.NullUUID{}, err
	} else if exists {
		return getExistedUserSession(database, info)
	} else {
		return registerNotExistedUser(database, info, emailConfig)
	}
}


func getExistedUserSession(database *gorm.DB, info structures.LoginInfo) (*db.Session, uuid.NullUUID, error) {
	existedUser, err := db.SelectUserByEmail(database, info.Email)

	if err != nil {
		return nil, uuid.NullUUID{}, err

	} else if existedUser.RegistrationID.Valid {
		return nil, existedUser.RegistrationID.UUID, nil

	} else {
		session, err := db.GetSession(database, existedUser.ID)
		return &session, uuid.NullUUID{}, err
	}
}


func registerNotExistedUser(database *gorm.DB, info structures.LoginInfo, emailConfig EmailConfiguration) (*db.Session, uuid.NullUUID, error) {
	var registrationID = uuid.NewV4()
	newUser := db.UserModel{uuid.NewV4(), info.Email, &registrationID}

	err := database.Create(&newUser).Error

	if err == nil {
		go sendVerificationEmail(info.Email, registrationID, emailConfig)
	}

	return nil, uuid.NullUUID{UUID: registrationID, Valid: true}, err
}


func sendVerificationEmail(toEmail string, registrationID uuid.UUID, emailConfig EmailConfiguration) {
	var link = "http://" + emailConfig.GetDomain() + "/v1/verify_register/" + registrationID.String()
	var signUpHTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+ "\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"

	utils.SendEmail(emailConfig.GetGmailLogin(), emailConfig.GetGmailPassword(), []string{toEmail}, "Verification", signUpHTML)
}

