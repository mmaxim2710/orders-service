package database

import (
	"github.com/mmaxim2710/orders-service/internal/repository/model"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	return nil
}
