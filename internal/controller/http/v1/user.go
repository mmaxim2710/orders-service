package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/usecase"
)

type userRoutes struct {
	u usecase.User
}

func newUserRoutes(handler fiber.Router, u usecase.User) {
	r := &userRoutes{
		u: u,
	}

	h := handler.Group("/user")
	{
		h.Post("/register", r.testHandle)
	}
}

func (u *userRoutes) testHandle(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
