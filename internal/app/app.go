package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mmaxim2710/orders-service/config"
	v1 "github.com/mmaxim2710/orders-service/internal/controller/http/v1"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"github.com/mmaxim2710/orders-service/internal/usecase"
	"github.com/mmaxim2710/orders-service/internal/usecase/repo/gormrepo"
	"github.com/mmaxim2710/orders-service/pkg/database"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"github.com/mmaxim2710/orders-service/pkg/validations"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	db, err := database.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - database.NewUserRepository: %w", err))
	}

	userRepo := gormrepo.NewUserRepository(db, l)
	tokenRepo := gormrepo.NewTokenRepository(db, l)
	serviceRepo := gormrepo.NewServiceRepo(db, l)

	// Use case
	userUseCase := usecase.NewUserUseCase(userRepo, tokenRepo)
	serviceUseCase := usecase.NewServiceUseCase(serviceRepo)

	// Validator
	validations.InitValidator()

	// JWT token init
	utils.InitJWTParams([]byte(cfg.JWT.Secret), *jwt.SigningMethodHS256)

	// HTTP server
	handler := fiber.New()
	v1.SetupRouter(handler, userUseCase, serviceUseCase, l)
	err = handler.Listen(cfg.Server.Port)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Lister: %w", err))
	}
}
