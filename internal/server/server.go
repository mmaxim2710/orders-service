package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/config"
	"gorm.io/gorm"
)

func NewServer(app *fiber.App, db *gorm.DB, config *config.Config) *Server {
	return &Server{
		App:    app,
		db:     db,
		Config: config,
	}
}
