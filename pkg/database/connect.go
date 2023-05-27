package database

import (
	"fmt"
	"github.com/mmaxim2710/orders-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func generateDSN(config *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Name, config.DB.SSLMode)
	return dsn
}

func New(cfg *config.Config) (*gorm.DB, error) {
	dsn := generateDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = Automigrate(db)
	if err != nil {
		return nil, err
	}

	DB = db

	return DB, nil
}
