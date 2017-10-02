package controller

import (
	"github.com/Nekitosss/authorization/db"
	"github.com/jinzhu/gorm"

	"github.com/satori/go.uuid"
)

func ValidateSession(database *gorm.DB, sessionID uuid.UUID) uuid.UUID {
	return db.ValidateSession(database, sessionID)
}

