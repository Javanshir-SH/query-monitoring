package http

import (
	"context"
	"fmt"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewServer(lc fx.Lifecycle, cfg *config.Config, logger *zap.Logger, db *gorm.DB) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			//logger.Info("Starting HTTP server...")
			go func() {
				err := app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
				if err != nil {
					logger.Error(err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			//logger.Info("Stopping HTTP server.")
			return app.Shutdown()
		},
	})

	return app
}
