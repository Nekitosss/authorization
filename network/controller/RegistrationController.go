package controller

import (
	"database/sql"

	"github.com/Nekitosss/authorization/db"

	"github.com/Nekitosss/authorization/network/controller/structures"
	"github.com/Nekitosss/authorization/utils"

	"github.com/satori/go.uuid"
)


type EmailConfiguration interface {

	GetGmailLogin() string

	GetGmailPassword() string

	GetDomain() string
}


//Регистрирует нового пользователя
func RegisterNewUser(database *sql.DB, info structures.RegisterInfo, emailConfig EmailConfiguration) error {
	
	filledInfo := fullFillRegisterInfo(info)
	
	alreadyExists, err := validateLoginExistance(database, filledInfo.Login)
	
	if err != nil {
		return err
	} else if alreadyExists {
		return utils.BLErrorUserAlreadyRegistered
	}
	
	return registerValidatedUser(database, filledInfo, emailConfig)
}

//Заполняет недостоящую информацию.
func fullFillRegisterInfo(info structures.RegisterInfo) structures.RegisterInfo {
	
	if info.Login == "" {
		info.Login = info.Email
	}
	
	if info.NameAlias == "" {
		info.NameAlias = info.Login
	}
	
	return info
}

func validateLoginExistance(database *sql.DB, login string) (bool, error) {
	return db.IsLoginExists(database, login)
}


func registerValidatedUser(database *sql.DB, info structures.RegisterInfo, emailConfig EmailConfiguration) error {
	
	var newHashedPassword, err = utils.CryptPassword([]byte(info.Password))
	
	if err != nil {
		return err
	}
	
	var registrationID = uuid.NewV4()
	newUser := db.Model{uuid.NewV4(), info.Login, info.Email, newHashedPassword, info.NameAlias, registrationID.String()}
	
	err = newUser.Insert(database)
	
	if err == nil {
		go sendVerification(info.Login, info.Email, info.RegisterType, registrationID, emailConfig)
	}
	
	return err
}


func sendVerification(login string, email string, verificationType structures.RegisterVerificationType, registrationID uuid.UUID, emailConfig EmailConfiguration) {
	
	switch verificationType {
	
	case structures.RegisterTypeEmail:
		sendVerificationEmail(email, registrationID, emailConfig)
	
	case structures.RegisterTypeTelegram:
		sendVerificationTelegramCode(login)
		
	}
	
}


func sendVerificationEmail(toEmail string, registrationID uuid.UUID, emailConfig EmailConfiguration) {

	var link = "http://" + emailConfig.GetDomain() + "/v1/verify_register/" + registrationID.String()
	var signUpHTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+ "\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"
	
	utils.SendEmail(emailConfig.GetGmailLogin(), emailConfig.GetGmailPassword(), []string{toEmail}, "Verification", signUpHTML)
}


func sendVerificationTelegramCode(login string) {

}