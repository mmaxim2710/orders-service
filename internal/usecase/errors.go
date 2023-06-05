package usecase

import "errors"

var (
	ErrUserExists               = errors.New("user with provided email is exists")
	ErrUserNotExists            = errors.New("user with provided email is not exists")
	ErrPasswordMismatch         = errors.New("password mismatch")
	ErrInvalidToken             = errors.New("invalid token")
	ErrServiceNotExists         = errors.New("service with provided params not exists")
	ErrUserHasNonClosedServices = errors.New("user has non closed services")
	ErrEmptySlice               = errors.New("empty slice")
)
