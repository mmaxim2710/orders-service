package service

import "errors"

var (
	ErrParseEnvConfig = errors.New("failed to parse config from env")
)
