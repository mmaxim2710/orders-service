package userrepository

import (
	"github.com/mmaxim2710/orders-service/internal/repository/model"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(user *model.User) (*model.User, error)
	FindByEmail(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) (*model.User, error)
	IsUserExists(email string) (bool, error)
}
