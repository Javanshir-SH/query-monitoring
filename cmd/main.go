package main

import (
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/config"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/db/pg"
	cache "github.com/Javanshir-SH/query-monitoring/internal/pkg/db/redis"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/logger"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/server/http"
	"github.com/Javanshir-SH/query-monitoring/internal/query/delivery/http/v1"
	"github.com/Javanshir-SH/query-monitoring/internal/query/service"
	"github.com/Javanshir-SH/query-monitoring/internal/query/storage"
	apiv1 "github.com/Javanshir-SH/query-monitoring/internal/todo/delivery/http/v1"
	svc "github.com/Javanshir-SH/query-monitoring/internal/todo/service"
	stg "github.com/Javanshir-SH/query-monitoring/internal/todo/storage"

	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			//query service dependencies
			config.NewConfig,
			logger.NewLogger,
			pg.NewGorm,
			cache.NewRedisClient,
			fx.Annotate(storage.NewQueryMonitoringRepo, fx.As(new(service.QueryStatementsRepository), new(error))),
			fx.Annotate(storage.NewQueryMonitoringCache, fx.As(new(service.CacheRepo), new(error))),
			fx.Annotate(service.NewQueryMonitoringService, fx.As(new(v1.Service), new(error))),
			v1.NewQueryMonitoringHandler,

			// todo service dependencies
			fx.Annotate(stg.NewTaskRepo, fx.As(new(svc.TaskRepository), new(error))),
			fx.Annotate(svc.NewTaskService, fx.As(new(apiv1.Service), new(error))),
			apiv1.NewTodoHandler,

			// server
			http.NewServer,
		),

		fx.Invoke(v1.RegisterQueryMonitoringHandler, apiv1.RegisterTodoHandler),
	).Run()
}
