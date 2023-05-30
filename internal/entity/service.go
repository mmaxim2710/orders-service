package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	User        User      `gorm:"foreignkey:UserID"`
	UserID      uuid.UUID `gorm:"type:uuid,not_null" json:"user_id"`
	Title       string    `gorm:"type:varchar(128)" json:"title"`
	Description string    `gorm:"type:varchar(1024)" json:"description"`
	Price       float64   `gorm:"type:numeric" json:"price"`
	IsClosed    bool      `gorm:"type:boolean" json:"is_closed"`
}
