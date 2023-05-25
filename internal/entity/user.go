package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"uuid" database:"type:uuid;primary_key;"`
	Login     string    `json:"login" database:"type:varchar(255)"`
	Email     string    `json:"email" database:"type:varchar(320)"`
	FirstName string    `json:"first_name" database:"type:varchar(128)"`
	LastName  string    `json:"last_name" database:"type:varchar(128)"`
	Password  string    `json:"password" database:"type:varchar(64)"`
}
