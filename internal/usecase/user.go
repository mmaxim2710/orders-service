package usecase

import "github.com/mmaxim2710/orders-service/internal/entity"

type UserUseCase struct {
	repo UserRepo
}

func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u UserUseCase) RegisterUser(login string, email string, firstName string, lastName string, password string) entity.Response {
	//TODO implement me
	panic("implement me")
}
