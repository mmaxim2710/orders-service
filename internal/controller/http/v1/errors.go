package v1

import "errors"

var (
	ErrValidationFailed = errors.New("validation failed")
	ErrEntitiesMismatch = errors.New("entities mismatch")
)
