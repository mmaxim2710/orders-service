package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/config"
	"github.com/mmaxim2710/orders-service/internal/pkg/database"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"gorm.io/gorm"
)

func newServer(app *fiber.App, db *gorm.DB, config *config.Config) *Server {
	return &Server{
		app:    app,
		db:     db,
		config: config,
	}
}

func Start() error {
	conf, err := config.GetConfig()
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}

	db, err := database.InitializeDB(conf.DB)
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}

	app := fiber.New()

	s := newServer(app, db, conf)
	s.SetupHandlers()
	return s.app.Listen(s.config.Server.Port)
}

func (s *Server) SetupHandlers() {
	orders := s.app.Group("/api")
	orders.Get("/ping", s.Ping)
}
