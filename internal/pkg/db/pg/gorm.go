package pg

import (
	"context"
	"fmt"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(lc fx.Lifecycle, conf *config.Config, logger *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.PgDB.Host,
		conf.PgDB.Port,
		conf.PgDB.User,
		conf.PgDB.Password,
		conf.PgDB.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			conn, err := db.DB()
			if err != nil {
				return err
			}

			err = conn.Close()
			if err != nil {
				logger.Error(err.Error())
			}

			return nil
		},
	})

	return db
}
