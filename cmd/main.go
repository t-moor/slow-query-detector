package main

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"

	"github.com/t-moor/slow-query-detector/internal/pkg/provider"
	"github.com/t-moor/slow-query-detector/internal/query/handler"
	"github.com/t-moor/slow-query-detector/internal/query/repository"
	"github.com/t-moor/slow-query-detector/internal/query/service"
)

func main() {
	app := fx.New(
		fx.Provide(provider.NewLogger),
		fx.Provide(provider.NewConfig),
		fx.Provide(provider.NewGorm),
		fx.Provide(validator.New),
		fx.Provide(repository.NewQueryAnalytics),
		fx.Provide(service.NewQueryAnalytics),
		fx.Provide(handler.NewQueryAnalytics),
		fx.Provide(func(analytics *service.QueryAnalytics) handler.Service {
			return analytics
		}),
		fx.Provide(func(analytics *repository.QueryAnalytics) service.Repository {
			return analytics
		}),
		fx.Invoke(provider.Server),
	)
	app.Run()
}
