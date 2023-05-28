package usecase

import "errors"

var (
	ErrUserExists       = errors.New("user with provided email is exists")
	ErrUserNotExists    = errors.New("user with provided email is not exists")
	ErrPasswordMismatch = errors.New("password mismatch")
	ErrInvalidToken     = errors.New("invalid token")
)
