package db

import (

	"github.com/satori/go.uuid"

)

type UserModel struct {
	
	ID uuid.UUID
	
	Email string `gorm:"index:user_email_idx;not null"`
	
	RegistrationID uuid.NullUUID `gorm:"index:user_registration_id_idx;"`
}

func (pc UserModel) TableName() string {
	return "auth.user_model_models"
}