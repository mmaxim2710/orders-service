package database

import (
	"github.com/mmaxim2710/orders-service/internal/entity"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&entity.Token{}); err != nil {
		return err
	}

	return nil
}
