package usecase

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"gorm.io/gorm"
)

type UserUseCase struct {
	userRepo  UserRepo
	tokenRepo TokenRepo
}

func New(u UserRepo, t TokenRepo) *UserUseCase {
	return &UserUseCase{
		userRepo:  u,
		tokenRepo: t,
	}
}

func (u UserUseCase) RegisterUser(login string, email string, firstName string, lastName string, password string) (*entity.User, error) {
	isUserExists, err := u.userRepo.IsUserExists(email)
	if isUserExists {
		return nil, ErrUserExists
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

	return u.userRepo.Create(user)
}

func (u UserUseCase) Login(email string, password string) (map[string]interface{}, error) {
	isExist, err := u.userRepo.IsUserExists(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if !isExist {
		return nil, ErrUserNotExists
	}

	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	ok := utils.ComparePassword(user.Password, password)
	if !ok {
		return nil, ErrPasswordMismatch
	}

	var rt string
	jwtToken, exp, err := utils.GenerateJWTToken(*user)
	if err != nil {
		return nil, err
	}

	token, err := u.tokenRepo.GetActiveToken(user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if token.Revoked == false {
		err = u.tokenRepo.Revoke(user.ID)
		if err != nil {
			return nil, err
		}
	}

	rt, _ = utils.GenerateRefreshToken()
	err = u.tokenRepo.Create(user.ID, rt)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user_id":       user.ID.String(),
		"token":         jwtToken,
		"refresh_token": rt,
		"exp":           exp,
	}, nil
}

func (u UserUseCase) Refresh(token string, userID string) (map[string]interface{}, error) {
	rt, err := u.tokenRepo.GetActiveToken(uuid.MustParse(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if rt.Token != token {
		return nil, ErrInvalidToken
	}

	err = u.tokenRepo.Revoke(uuid.MustParse(userID))
	if err != nil {
		return nil, err
	}

	newrt, _ := utils.GenerateRefreshToken()
	err = u.tokenRepo.Create(uuid.MustParse(userID), newrt)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	jwttoken, exp, err := utils.GenerateJWTToken(*user)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user_id":       user.ID.String(),
		"token":         jwttoken,
		"refresh_token": newrt,
		"exp":           exp,
	}, nil
}

func (u UserUseCase) Update(userID uuid.UUID, email string, firstName string, lastName string) (*entity.User, error) {
	isExist, err := u.userRepo.IsUserExists(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if isExist {
		return nil, ErrUserExists
	}

	newUser := &entity.User{
		ID:        userID,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	updatedUser, err := u.userRepo.Update(newUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
