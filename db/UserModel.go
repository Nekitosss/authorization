package db

import (

	"github.com/satori/go.uuid"

)

type UserModel struct {
	
	ID uuid.UUID
	
	Login string `gorm:"index:user_login_idx;not null"`
	
	Email string `gorm:"index:user_email_idx;not null"`
	
	PasswordHash []byte `gorm:"not null"`
	
	NameAlias string `gorm:"not null"`
	
	RegistrationID uuid.UUID `gorm:"index:user_registration_id_idx;"`
}

func (pc UserModel) TableName() string {
	return "auth.user_model_models"
}