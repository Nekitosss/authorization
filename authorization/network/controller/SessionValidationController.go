package controller

import (
	"authorization/db"
	"database/sql"

	"github.com/satori/go.uuid"
)

func ValidateSession(database *sql.DB, sessionID uuid.UUID) uuid.UUID {
	return db.ValidateSession(database, sessionID)
}

