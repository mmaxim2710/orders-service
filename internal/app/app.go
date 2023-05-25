package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/config"
	v1 "github.com/mmaxim2710/orders-service/internal/controller/http/v1"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"github.com/mmaxim2710/orders-service/internal/usecase"
	"github.com/mmaxim2710/orders-service/internal/usecase/repo"
	"github.com/mmaxim2710/orders-service/pkg/database"
)

func Run(cfg *config.Config) {
	// Repository
	db, err := database.New(cfg)
	if err != nil {
		utils.Logger.Error(fmt.Errorf("app - Run - database.New: %w", err).Error())
	}

	userRepo := repo.New(db)

	// Use case
	userUseCase := usecase.New(userRepo)

	// HTTP server
	handler := fiber.New()
	v1.SetupRouter(handler, userUseCase)
}