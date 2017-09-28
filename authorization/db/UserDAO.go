package db

import (
	"database/sql"

	"github.com/satori/go.uuid"
)


func (u Model) Insert(database *sql.DB) error {
	var _, err = database.Exec(insertSQL, u.ID.String(), u.Login, u.Email, string(u.PasswordHash), u.NameAlias, u.RegistrationID)

	return err
}

func (u Model) UpdateRegistrationID(database *sql.DB, newRegistrationID string) error {
	var _, err = database.Exec(updateSQL, newRegistrationID, u.ID.String())

	return err
}

func SelectUserByLogin(database *sql.DB, login string) (Model, error) {
	return selectUser(database, selectUserByLoginSQL, login)
}


func SelectUserByRegisterID(database *sql.DB, registerID uuid.UUID) (Model, error) {
	return selectUser(database, selectUseByRegIDSQL, registerID.String())
}


func selectUser(database *sql.DB, sql string, arg interface{}) (Model, error) {
	var m Model
	var passwordHash []byte
	
	var err = database.QueryRow(sql, arg).Scan(&m.ID, &m.Login, &m.Email, &passwordHash, &m.NameAlias, &m.RegistrationID)
	
	if err != nil {
		return m, err
	}
	
	m.PasswordHash = passwordHash
	
	return m, nil
	
}


func GetUserPasswordHash(database *sql.DB, login string) ([]byte, error) {
	var model, err = SelectUserByLogin(database, login)

	return model.PasswordHash, err
}

func IsLoginExists(database *sql.DB, login string) (bool, error) {
	var exists bool
	var err = database.QueryRow(existsLoginValidationSQL, login).Scan(&exists)

	return exists, err
}
