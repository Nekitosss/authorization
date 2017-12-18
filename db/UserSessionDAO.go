package db

import (
	"github.com/jinzhu/gorm"
	"time"

	"github.com/Nekitosss/authorization/utils"

	"github.com/satori/go.uuid"
)



func GetSession(database  *gorm.DB, userid uuid.UUID) (Session, error) {

	var session = Session{}

	err := database.Where("user_model_id = $1", userid).Find(&session).Error

	if err != gorm.ErrRecordNotFound {
		return session, err

	} else if err == gorm.ErrRecordNotFound {
		return createSession(database, userid)

	} else {
		return session, nil
	}

}


func ValidateSession(database *gorm.DB, sessionID uuid.UUID) uuid.UUID {

	var session = Session{}
	session.SessionID = sessionID

	err := database.First(&session).Error

	if err != gorm.ErrRecordNotFound {
		utils.CheckError(err)
	}

	return session.UserModelID
}


func createSession(database *gorm.DB, userid uuid.UUID) (Session, error) {

	var session = Session{}
	session.SessionID = uuid.NewV4()
	session.UserModelID = userid
	session.LoginTime = time.Now()
	session.LastSeenTime = session.LoginTime

	err := database.Create(&session).Error

	return session, err
}