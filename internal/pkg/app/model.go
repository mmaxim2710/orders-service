package app

import (
	"github.com/mmaxim2710/orders-service/internal/endpoint"
	"github.com/mmaxim2710/orders-service/internal/server"
	"github.com/mmaxim2710/orders-service/internal/service"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	server   *server.Server
}
