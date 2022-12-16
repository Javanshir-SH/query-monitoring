package redis

import (
	"context"
	"fmt"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewRedisClient(lc fx.Lifecycle, cfg *config.Config, logger *zap.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.DB,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Checking redis readiness...")
			if _, err := client.Ping(context.Background()).Result(); err != nil {
				logger.Error(err.Error())
			}
			logger.Info("redis connection pool ready to use.")
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Info("Closing redis connection pool...")
			err := client.Close()
			if err != nil {
				logger.Error(err.Error())
			}
			return nil
		},
	})

	return client
}
