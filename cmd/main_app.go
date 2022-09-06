package cmd

import (
	"context"

	"github.com/kaankoken/helper/pkg"
	"github.com/kaankoken/helper/pkg/helper"
	"go.uber.org/fx"
)

// MainApp -> Creates new app that using all DI modules
func MainApp() *fx.App {
	app := fx.New(
		pkg.FakeTagModule,
		pkg.FakeFlavorDevModule,
		helper.LoggerModule,
		fx.Invoke(RegisterHooks),
	)

	return app
}

// RegisterHooks -> Registering lifecycle of fx & running http server (Gin)
func RegisterHooks(lifecycle fx.Lifecycle, logger *helper.LogHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Starting application")
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping application")
				return nil
			},
		},
	)
}
