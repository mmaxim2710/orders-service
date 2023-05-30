package usecase

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
)

type ServiceUseCase struct {
	serviceRepo ServiceRepo
}

func NewServiceUseCase(repo ServiceRepo) *ServiceUseCase {
	return &ServiceUseCase{
		serviceRepo: repo,
	}
}

func (s *ServiceUseCase) Create(userID uuid.UUID, title string, description string, price float64) (*entity.Service, error) {
	newService := &entity.Service{
		UserID:      userID,
		Title:       title,
		Description: description,
		Price:       price,
		IsClosed:    false,
	}

	return s.serviceRepo.Create(newService)
}
