package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	App    *fiber.App
	db     *gorm.DB
	Config *config.Config
}
