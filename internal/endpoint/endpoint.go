package endpoint

import (
	"github.com/gofiber/fiber/v2"
)

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString(e.s.HelloWorld())
}
