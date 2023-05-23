package utils

import "go.uber.org/zap"

var (
	Logger *zap.Logger
)

func InitLogger() error {
	if Logger != nil {
		logger, err := zap.NewProduction()
		Logger = logger
		if err != nil {
			return err
		}
	}

	return nil
}
