package usecase

import (
	"errors"
	"fmt"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
)

type UserUseCase struct {
	repo UserRepo
}

func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u UserUseCase) RegisterUser(login string, email string, firstName string, lastName string, password string) (*entity.User, error) {
	isUserExists, err := u.repo.IsUserExists(email)
	if isUserExists {
		return nil, errors.New(fmt.Sprintf("user with email %s is exists in DB", email))
	}

	encrPassword, err := utils.EncryptString(password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Login:     login,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  encrPassword,
	}

	return u.repo.Create(user)
}
