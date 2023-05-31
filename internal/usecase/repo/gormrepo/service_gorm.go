package gormrepo

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
	l  logger.Interface
}

func NewServiceRepo(db *gorm.DB, l logger.Interface) *ServiceRepository {
	return &ServiceRepository{
		db: db,
		l:  l,
	}
}

func (s *ServiceRepository) Create(service *entity.Service) (*entity.Service, error) {
	s.l.Info("serviceRepo - Create: Creating new service: title=%s, desc=%s, price=%f",
		service.Title, service.Description, service.Price)
	newService := &entity.Service{}
	err := s.db.Model(&entity.Service{}).
		Create(&service).
		First(&newService).Error
	return newService, err
}

func (s *ServiceRepository) Update(service *entity.Service) (*entity.Service, error) {
	return nil, nil
}

func (s *ServiceRepository) Delete(service *entity.Service) (*entity.Service, error) {
	return nil, nil
}

func (s *ServiceRepository) GetServicesByUserId(userID uuid.UUID) (*entity.Service, error) {
	return nil, nil
}
