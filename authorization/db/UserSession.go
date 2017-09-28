package db

import (
	"time"

	"github.com/satori/go.uuid"
)

type Session struct {
	SessionID uuid.UUID

	UserID uuid.UUID

	LoginTime time.Time

	LastSeenTime time.Time
}
