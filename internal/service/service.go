package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"github.com/mmaxim2710/orders-service/internal/repository/model"
	"github.com/mmaxim2710/orders-service/internal/repository/userrepository"
	"gorm.io/gorm"
)

func New(userRepo userrepository.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) HelloWorld() string {
	return "Hello world!"
}

func (s *Service) RegisterUser(login string, email string, firstName string, lastName string, password string) Response {
	isExists, err := s.userRepo.IsUserExists(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      err.Error(),
		}
	}

	if isExists {
		return Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "User already exists",
			Data:       nil,
			Error:      "User already exists",
		}
	}

	encPass, err := utils.EncryptString(password)
	if err != nil {
		return Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      err.Error(),
		}
	}

	user, err := s.userRepo.Create(&model.User{
		Login:     login,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  encPass,
	})

	if err != nil {
		return Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      err.Error(),
		}
	}

	return Response{
		StatusCode: fiber.StatusOK,
		Message:    "User created",
		Data: map[string]string{
			"id":         user.ID.String(),
			"login":      user.Login,
			"email":      user.Email,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
		Error: nil,
	}
}
