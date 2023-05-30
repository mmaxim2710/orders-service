package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
)

func Protected() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   utils.JwtSignatureKey,
		ErrorHandler: jwtError,
		ContextKey:   "jwt",
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"StatusCode": fiber.StatusBadRequest,
			"Message":    "Missing or malformed JWT",
			"Data":       nil,
			"Error":      err,
		})
	} else {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"StatusCode": fiber.StatusUnauthorized,
			"Message":    "Invalid or expired JWT",
			"Data":       nil,
			"Error":      err,
		})
	}
}
