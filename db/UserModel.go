package db

import (

	"github.com/satori/go.uuid"

)

type UserModel struct {
	
	ID uuid.UUID
	
	Email string `gorm:"index:user_email_idx;not null"`

}


func (pc UserModel) TableName() string {
	return "auth.user_model_models"
}


type UserRegistration struct {

	UserID uuid.UUID `gorm:"primary_key"`

	RegistrationID uuid.UUID `gorm:"primary_key;"`

	Confirmed bool

}


func (ur UserRegistration) TableName() string {
	return "auth.user_registration_models"
}