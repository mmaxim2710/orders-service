package usecase

import (
	"github.com/mmaxim2710/orders-service/internal/entity"
)

type (
	User interface {
		RegisterUser(login string, email string, firstName string, lastName string, password string) (*entity.User, error)
	}

	UserRepo interface {
		Create(user *entity.User) (*entity.User, error)
		FindByID(user *entity.User) (*entity.User, error)
		FindByEmail(user *entity.User) (*entity.User, error)
		Update(user *entity.User) (*entity.User, error)
		Delete(user *entity.User) (*entity.User, error)
		IsUserExists(email string) (bool, error)
	}
)
