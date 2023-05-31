package usecase

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
)

type (
	User interface {
		RegisterUser(login string, email string, firstName string, lastName string, password string) (*entity.User, error)
		Login(email string, password string) (map[string]interface{}, error)
		Refresh(token string, userID string) (map[string]interface{}, error)
		Update(userID uuid.UUID, email string, firstName string, lastName string) (*entity.User, error)
	}

	UserRepo interface {
		Create(user *entity.User) (*entity.User, error)
		FindByID(id string) (*entity.User, error)
		FindByEmail(email string) (*entity.User, error)
		Update(user *entity.User) (*entity.User, error)
		Delete(user *entity.User) (*entity.User, error)
		IsUserExists(email string) (bool, error)
	}
)

type (
	TokenRepo interface {
		Create(userID uuid.UUID, token string) error
		GetActiveToken(userID uuid.UUID) (entity.Token, error)
		Revoke(userID uuid.UUID) error
	}
)

type (
	Service interface {
		Create(userID uuid.UUID, title string, description string, price float64) (*entity.Service, error)
		Update(serviceID uuid.UUID, title string, description string, price float64, isClosed bool) (*entity.Service, error)
		GetByID(serviceID uuid.UUID) (*entity.Service, error)
	}
	ServiceRepo interface {
		Create(service *entity.Service) (*entity.Service, error)
		Update(service *entity.Service) (*entity.Service, error)
		Delete(service *entity.Service) (*entity.Service, error)
		IsServiceExists(serviceID uuid.UUID) (bool, error)
		GetServiceByID(userID uuid.UUID) (*entity.Service, error)
		GetServicesByUserID(userID uuid.UUID) (*entity.Service, error)
	}
)
