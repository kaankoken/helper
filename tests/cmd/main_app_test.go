package main

import (
	"context"
	"testing"

	"github.com/kaankoken/helper/cmd"
	"github.com/kaankoken/helper/pkg"
	"github.com/kaankoken/helper/pkg/helper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestMainApp(t *testing.T) {
	t.Run("test-main-app-func", func(t *testing.T) {
		app := fxtest.New(
			t,
			pkg.FakeFlavorDevModule,
			pkg.FakeTagModule,
			helper.LoggerModule,
			fx.Invoke(cmd.RegisterHooks),
		)

		app.RequireStart()
		defer app.RequireStop()
	})

	t.Run("test-main-app", func(t *testing.T) {
		app := cmd.MainApp()

		app.Start(context.Background())

		defer app.Stop(context.Background())
	})
}
