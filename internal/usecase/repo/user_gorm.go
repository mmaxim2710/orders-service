package repo

import (
	"github.com/mmaxim2710/orders-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *UserRepository) FindByID(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) FindByEmail(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) Update(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) Delete(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) IsUserExists(email string) (bool, error) {
	var count int64
	err := u.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
