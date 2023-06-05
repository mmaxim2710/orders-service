package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int       `gorm:"primary_key;unique;not_null;auto_increment"`
	OrderID   uuid.UUID `gorm:"type:uuid;not_null" json:"id"`
	Service   Service   `gorm:"foreignkey:ServiceID"`
	ServiceID uuid.UUID `gorm:"type:uuid;not_null" json:"service_id"`
	Status    string    `gorm:"type:varchar(32);not_null"`
}
