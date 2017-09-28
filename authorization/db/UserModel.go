package db

import "github.com/satori/go.uuid"

type Model struct {
	
	ID uuid.UUID
	
	Login string
	
	Email string
	
	PasswordHash []byte
	
	NameAlias string
	
	RegistrationID string
}

