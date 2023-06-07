package gormrepo

import (
	"github.com/google/uuid"
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
	u.l.Info("userRepo - FindByID: Finding user by id %s", id)
	var user entity.User
	err := u.db.Limit(1).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	u.l.Info("userRepo - FindByEmail: Finding user by email %s", email)
	var user entity.User
	err := u.db.Limit(1).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) Update(user *entity.User) (*entity.User, error) {
	u.l.Info("userRepo - Update: Updating user %v", user.ID)
	var newUser entity.User
	err := u.db.Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(user).
		First(&newUser).Error
	return &newUser, err
}

func (u *UserRepository) Delete(userID uuid.UUID) (*entity.User, error) {
	u.l.Info("userRepo - Delete: Deleting user with id %s", userID.String())
	var delUser entity.User
	err := u.db.Model(&entity.User{}).
		Where("id = ?", userID).
		First(&delUser).
		Delete(&entity.User{}).Error
	return &delUser, err
}

func (u *UserRepository) IsUserExistsByEmail(email string) (bool, error) {
	u.l.Info("userRepo - IsUserExistsByEmail: Checking if user with email %s is exists", email)
	var count int64
	err := u.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		u.l.Error(err, "userRepo - IsUserExistsByEmail")
	}
	return count > 0, err
}

func (u *UserRepository) IsUserExistsByUserID(userID uuid.UUID) (bool, error) {
	u.l.Info("userRepo - IsUserExistsByUserID: Checking if user with id %v is exists", userID)
	var count int64
	err := u.db.Model(&entity.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		u.l.Error(err, "userRepo - IsUserExistsByUserID")
	}
	return count > 0, err
}
