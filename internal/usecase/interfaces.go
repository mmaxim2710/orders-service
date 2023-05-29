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

	TokenRepo interface {
		Create(userID uuid.UUID, token string) error
		GetActiveToken(userID uuid.UUID) (entity.Token, error)
		Revoke(userID uuid.UUID) error
	}
)
