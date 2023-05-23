package service

import (
	_ "github.com/caarlos0/env/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	app    *fiber.App
	db     *gorm.DB
	config *config.Config
}
