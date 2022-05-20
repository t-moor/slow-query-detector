package provider

import (
	"context"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/t-moor/slow-query-detector/internal/query/handler"
)

func Server(lc fx.Lifecycle, logger *zap.Logger, queryHandler *handler.QueryAnalytics, conf *viper.Viper) {
	srv := fiber.New()
	srv.Get("/queries", queryHandler.FindQueries)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server")
			go func() {
				err := srv.Listen(":" + conf.GetString("port"))
				if err != nil {
					logger.Error(err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server")
			return srv.Shutdown()
		},
	})
}
