package userrepository

import (
	"github.com/mmaxim2710/orders-service/internal/repository/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user *model.User) (*model.User, error) {
	return user, u.db.Create(&user).Error
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

func (u *userRepository) IsUserExists(email string) (bool, error) {
	var count int64
	err := u.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
