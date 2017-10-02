package db

import (
	"time"

	"github.com/satori/go.uuid"
)

type Session struct {
	SessionID uuid.UUID `gorm:"primary_key"`

	UserModelID uuid.UUID `gorm:"index:session_user_id_idx;not null"`

	LoginTime time.Time `gorm:"index:session_user_id_idx;not null"`

	LastSeenTime time.Time `gorm:"index:session_user_id_idx;not null"`
}


func (pc Session) TableName() string {
	return "auth.session"
}