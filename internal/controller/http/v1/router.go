package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/usecase"
)

func SetupRouter(handler *fiber.App, u usecase.User) {
	handler.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	h := handler.Group("/v1")
	{
		newUserRoutes(h, u)
	}
}
