package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/service"
	"log"
)

func main() {
	c, err := service.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(c.Server.Port))
}
