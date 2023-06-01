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
	s.l.Info("serviceRepo - Update: Updating service with values: title=%s, desc=%s, price=%f",
		service.Title, service.Description, service.Price)
	var newService entity.Service
	err := s.db.Model(&entity.Service{}).
		Where("id = ?", service.ID).
		Updates(service).
		First(&newService).Error
	return &newService, err
}

func (s *ServiceRepository) Delete(service *entity.Service) (*entity.Service, error) {
	return nil, nil
}

func (s *ServiceRepository) IsServiceExists(serviceID uuid.UUID) (bool, error) {
	s.l.Info("serviceRepo - IsServiceExists: Checking if service with id=%s is exists", serviceID.String())
	var count int64
	err := s.db.Model(&entity.Service{}).Where("id = ?", serviceID).Count(&count).Error
	if err != nil {
		s.l.Error(err, "serviceRepo - IsServiceExists")
	}
	return count > 0, err
}

func (s *ServiceRepository) GetServiceByID(serviceID uuid.UUID) (*entity.Service, error) {
	s.l.Info("serviceRepo - GetServiceByID: Getting service by id %s", serviceID.String())
	service := &entity.Service{}
	err := s.db.Model(&entity.Service{}).
		Where("id = ?", serviceID).
		First(service).Error
	return service, err
}

func (s *ServiceRepository) GetServicesByUserID(userID uuid.UUID) ([]entity.Service, error) {
	var services []entity.Service
	err := s.db.Model(&entity.Service{}).
		Where("user_id = ?", userID).
		Find(&services).Error
	return services, err
}
