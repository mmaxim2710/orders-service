package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/config"
)

func NewServer(app *fiber.App, config *config.Config) *Server {
	return &Server{
		App:    app,
		Config: config,
	}
}
