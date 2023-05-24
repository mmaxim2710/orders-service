package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/config"
	"github.com/mmaxim2710/orders-service/internal/endpoint"
	"github.com/mmaxim2710/orders-service/internal/pkg/database"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"github.com/mmaxim2710/orders-service/internal/repository/userrepository"
	"github.com/mmaxim2710/orders-service/internal/server"
	"github.com/mmaxim2710/orders-service/internal/service"
)

func New() (*App, error) {
	a := &App{}

	conf, err := config.GetConfig()
	if err != nil {
		utils.Logger.Error(err.Error())
		return nil, err
	}

	db, err := database.InitializeDB(conf.DB)
	if err != nil {
		utils.Logger.Error(err.Error())
		return nil, err
	}

	userRepo := userrepository.New(db)

	a.service = service.New(userRepo)
	a.endpoint = endpoint.New(a.service)

	app := fiber.New()

	s := server.NewServer(app, conf)
	a.server = s

	a.SetupHandlers()

	return a, nil
}

func (a *App) Run() error {
	return a.server.App.Listen(a.server.Config.Server.Port)
}

func (a *App) SetupHandlers() {
	orders := a.server.App.Group("/api/v1")
	orders.Get("/ping", a.endpoint.Ping)

	user := a.server.App.Group("/api/v1/user")
	user.Post("/register")
}
