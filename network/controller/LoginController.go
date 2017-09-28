package controller

import (
	"github.com/Nekitosss/authorization/db"
	"github.com/Nekitosss/authorization/network/controller/structures"
	"github.com/Nekitosss/authorization/utils"
	"database/sql"
)

func LogIn(database *sql.DB, info structures.LoginInfo) (*db.Session, error) {

	isCorrectLoginPassword, err := checkLoginAndPasswordEquality(database, info.Login, info.Password)

	if err != nil || isCorrectLoginPassword == false {
		return nil, utils.BLErrorLoginIncorrectPassword
	}

	user, err := db.SelectUserByLogin(database, info.Login)
	utils.CheckError(err)

	session, err := db.GetSession(database, user.ID)
	utils.CheckError(err)

	return &session, nil
}



func checkLoginAndPasswordEquality(database *sql.DB, login string, password string) (bool, error) {

	var dbPassword, err = db.GetUserPasswordHash(database, login)

	//TODO: Доделать. Проверить конкретный тип ошибки
	if err != nil {
		return false, err
	}

	var equals = utils.CompareHashAndPassword(dbPassword, []byte(password))

	return equals == nil, nil

}
