package usecase

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"gorm.io/gorm"
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

func (s *ServiceUseCase) Update(serviceID uuid.UUID, title string, description string, price float64, isClosed bool) (*entity.Service, error) {
	isExists, err := s.serviceRepo.IsServiceExists(serviceID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if !isExists {
		return nil, ErrServiceNotExists
	}

	updService := &entity.Service{
		ID:          serviceID,
		Title:       title,
		Description: description,
		Price:       price,
		IsClosed:    isClosed,
	}

	return s.serviceRepo.Update(updService)
}

func (s *ServiceUseCase) GetByID(serviceID uuid.UUID) (*entity.Service, error) {
	isExists, err := s.serviceRepo.IsServiceExists(serviceID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if !isExists {
		return nil, ErrServiceNotExists
	}

	return s.serviceRepo.GetServiceByID(serviceID)
}
