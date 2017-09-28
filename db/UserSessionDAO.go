package db

import (
	"database/sql"
	"time"

	"github.com/Nekitosss/authorization/utils"

	"github.com/satori/go.uuid"
)

const (
	insertSessionSQL = "INSERT INTO" + sessionTable + "(" + sessionColumns + ") VALUES ($1, $2, $3, $4)"

	selectSessionSQL = "SELECT" + sessionColumns + "FROM" + sessionTable+  "WHERE userid = $1"

	sessionColumns = " identifier, userid, logintime, lastseentime "

	sessionTable = " auth.session "

	validateSessionSQL = "SELECT userid FROM " + sessionTable + " WHERE session = $1"
)

func (s Session) Insert(database *sql.DB) error {
	var _, err = database.Exec(insertSessionSQL, s.SessionID.String(), s.UserID.String(), s.LoginTime, s.LastSeenTime)
	return err
}


func GetSession(database  *sql.DB, userid uuid.UUID) (Session, error) {

	var session = Session{}

	var err = database.QueryRow(selectSessionSQL, userid).Scan(&session.SessionID, &session.UserID, &session.LoginTime, &session.LastSeenTime)

	if err != sql.ErrNoRows {
		return session, err

	} else if err == sql.ErrNoRows {

		return createSession(database, userid)
	} else {

		return session, nil
	}

}


func ValidateSession(database *sql.DB, sessionID uuid.UUID) uuid.UUID {

	userID := uuid.Nil

	var err = database.QueryRow(validateSessionSQL, sessionID).Scan(&userID)

	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return userID
}


func createSession(database *sql.DB, userid uuid.UUID) (Session, error) {

	var session = Session{}
	session.SessionID = uuid.NewV4()
	session.UserID = userid
	session.LoginTime = time.Now()
	session.LastSeenTime = session.LoginTime

	var err = session.Insert(database)

	return session, err
}