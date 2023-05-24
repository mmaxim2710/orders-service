package service

import (
	_ "github.com/caarlos0/env/v8"
	"github.com/mmaxim2710/orders-service/internal/repository/userrepository"
)

type Service struct {
	userRepo userrepository.UserRepository
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
