package network

import (
	"github.com/jinzhu/gorm"

	"github.com/Nekitosss/authorization/network/controller"
	"github.com/Nekitosss/authorization/network/controller/structures"

	"github.com/satori/go.uuid"

	"github.com/Nekitosss/authorization/db"
)


var executor Executor


func GetExecutor() Executor {
	return executor
}


func PrepareExecutor(database *gorm.DB, emailConfig EmailConfiguration) {
	executor = &ExecutorImpl{database, emailConfig}
}


type EmailConfiguration interface {

	GetGmailLogin() string

	GetGmailPassword() string

	GetDomain() string
}


type Executor interface {

	AuthUser(info structures.AuthInfo)  (*db.Session, error)

	VerifyRegister(registerID uuid.UUID) error

	Login(info structures.LoginInfo) (uuid.UUID, error)

	ValidateSession(info structures.ValidateSessionInfo) (uuid.UUID, error)

}


type ExecutorImpl struct {

	database *gorm.DB

	emailConfig EmailConfiguration

}


func (e *ExecutorImpl) AuthUser(info structures.AuthInfo)  (*db.Session, error) {
	return controller.AuthVerifiedUser(e.database, info)
}


func (e *ExecutorImpl) VerifyRegister(registerID uuid.UUID) error {
	return controller.VerifyRegistration(e.database, registerID)
}


func (e *ExecutorImpl) Login(info structures.LoginInfo) (uuid.UUID, error) {
	return controller.LogIn(e.database, info, e.emailConfig)
}


func (e *ExecutorImpl) ValidateSession(info structures.ValidateSessionInfo) (uuid.UUID, error) {
	sid, err := uuid.FromString(info.SessionID)
	if err != nil {
		return uuid.Nil, err
	}

	return controller.ValidateSession(e.database, sid), nil
}