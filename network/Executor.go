package network

import (
	"database/sql"

	"github.com/Nekitosss/authorization/network/controller"
	"github.com/Nekitosss/authorization/network/controller/structures"

	"github.com/satori/go.uuid"

	"github.com/Nekitosss/authorization/db"
)


var executor Executor


func PrepareExecutor(database *sql.DB, emailConfig EmailConfiguration) {
	executor = &ExecutorImpl{database, emailConfig}
}


type EmailConfiguration interface {

	GetGmailLogin() string

	GetGmailPassword() string

	GetDomain() string
}


type Executor interface {

	Register(info structures.RegisterInfo) error

	VerifyRegister(registerID uuid.UUID) error

	Login(info structures.LoginInfo) (*db.Session, error)

	ValidateSession(info structures.ValidateSessionInfo) (uuid.UUID, error)
}


type ExecutorImpl struct {
	database *sql.DB

	emailConfig EmailConfiguration
}


func (e *ExecutorImpl) Register(info structures.RegisterInfo) error {
	return controller.RegisterNewUser(e.database, info, e.emailConfig)
}

func (e *ExecutorImpl) VerifyRegister(registerID uuid.UUID) error {
	return controller.VerifyRegistration(e.database, registerID)
}

func (e *ExecutorImpl) Login(info structures.LoginInfo) (*db.Session, error) {
	return controller.LogIn(e.database, info)
}


func (e *ExecutorImpl) ValidateSession(info structures.ValidateSessionInfo) (uuid.UUID, error) {

	sid, err := uuid.FromString(info.SessionID)

	if err != nil {
		return uuid.Nil, err
	}

	return controller.ValidateSession(e.database, sid), nil
}