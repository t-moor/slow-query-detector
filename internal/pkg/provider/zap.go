package provider

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle) *zap.Logger {
	logger, _ := zap.NewProduction()

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			err := logger.Sync()
			if err != nil {
				logger.Error(err.Error())
			}

			return nil
		},
	})

	return logger
}
