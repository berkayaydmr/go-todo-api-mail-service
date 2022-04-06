package common

import "go.uber.org/zap"

func Logger(debug string) *zap.Logger {
	if debug == "true" {
		logger, _ := zap.NewDevelopment()
		return logger
	} else {
		logger, _ := zap.NewProduction()
		return logger
	}
}
