package database

import (
	"fmt"
	"github.com/mmaxim2710/orders-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(dbConfig config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = Automigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
