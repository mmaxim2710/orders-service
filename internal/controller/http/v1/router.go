package v1

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mmaxim2710/orders-service/docs/order-service"
	"github.com/mmaxim2710/orders-service/internal/usecase"
	"github.com/mmaxim2710/orders-service/pkg/logger"
)

// SetupRouter -.
// Swagger spec:
// @title       Orders Service
// @description Provides API for ordering
// @version     1.0
// @host        localhost:3000
// @BasePath    /v1
func SetupRouter(handler *fiber.App, u usecase.User, l logger.Interface) {
	handler.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	handler.Get("/swagger/*", swagger.HandlerDefault)

	h := handler.Group("/v1")
	{
		newUserRoutes(h, u, l)
	}
}
