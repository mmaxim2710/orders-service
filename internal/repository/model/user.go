package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key;"`
	Login     string    `json:"login" gorm:"type:varchar(255);not_null"`
	Email     string    `json:"email" gorm:"type:varchar(320);not_null;"`
	FirstName string    `json:"first_name" gorm:"type:varchar(128);not_null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(128);not_null"`
	Password  string    `json:"password" gorm:"not_null"`
}
