package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle) *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

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
