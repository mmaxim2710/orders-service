package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Login     string    `gorm:"type:varchar(255)" json:"login"`
	Email     string    `gorm:"type:varchar(320)" json:"email"`
	FirstName string    `gorm:"type:varchar(128)" json:"first_name"`
	LastName  string    `gorm:"type:varchar(128)" json:"last_name"`
	Password  string    `gorm:"type:varchar(64)" json:"password"`
}
