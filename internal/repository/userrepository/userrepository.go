package userrepository

import (
	"github.com/mmaxim2710/orders-service/internal/repository/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Create(user *model.User) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) FindByID(user *model.User) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) FindByEmail(user *model.User) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) Update(user *model.User) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) Delete(user *model.User) (*model.User, error) {
	return nil, nil
}
