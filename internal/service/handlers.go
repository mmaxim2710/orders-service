package service

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Ping(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
