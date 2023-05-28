package repo

import (
	"github.com/mmaxim2710/orders-service/internal/entity"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	l  logger.Interface
}

func NewUserRepository(db *gorm.DB, l logger.Interface) *UserRepository {
	return &UserRepository{
		db: db,
		l:  l,
	}
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	u.l.Info("userRepo - Create: Creating user")
	return user, u.db.Create(&user).Error
}

func (u *UserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
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
	if err != nil {
		u.l.Error(err, "userRepo - IsUserExists")
	}
	return count > 0, err
}
