package pkg_test

import (
	"testing"

	"github.com/kaankoken/helper/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestFlavor(t *testing.T) {
	t.Run("flavor=invoke-method", func(t *testing.T) {
		flavor := pkg.FakeFlavorDevConstructorForDi()

		assert.NotNil(t, flavor)
		assert.NotNil(t, &flavor)
	})
}

func TestFlavorWithFx(t *testing.T) {
	t.Run("flavor-dev=injection-test", func(t *testing.T) {
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			fx.Populate(&g),
			pkg.FakeFlavorDevModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("flavor-dev=injection-test-with-functions", func(t *testing.T) {
		var g fx.DotGraph
		var l *pkg.Flavor

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			fx.Populate(&g),
			fx.Populate(&l),
			pkg.FakeFlavorDevModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("flavor-release=injection-test", func(t *testing.T) {
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			fx.Populate(&g),
			pkg.FakeFlavorReleaseModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("flavor-release=injection-test-with-functions", func(t *testing.T) {
		var g fx.DotGraph
		var l *pkg.Flavor

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			fx.Populate(&g),
			fx.Populate(&l),
			pkg.FakeFlavorReleaseModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})
}
