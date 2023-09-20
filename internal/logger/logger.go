package logger

import (
	"go.uber.org/zap"
	"projects/grafit_info/config"
)

// NewLogger - service logger.
func NewLogger(cfg *config.Config) *zap.Logger {
	logger, _ := zap.NewProduction()

	if cfg.Env == "dev" {
		logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync()

	return logger
}
