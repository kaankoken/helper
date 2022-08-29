package pkg_test

import (
	"testing"

	"github.com/kaankoken/helper/pkg"
	"github.com/kaankoken/helper/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestTag(t *testing.T) {
	t.Run("tag=invoke-method", func(t *testing.T) {
		tag := pkg.FakeConstructorForDi()

		assert.NotNil(t, tag)
		assert.NotNil(t, &tag)
	})
}

func TestTagWithFx(t *testing.T) {
	t.Run("tag-=injection-test", func(t *testing.T) {
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			helper.DebugModule,
			fx.Populate(&g),
			pkg.FakeModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("d-logger=injection-test-with-functions", func(t *testing.T) {
		var g fx.DotGraph
		var l *pkg.Tag

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			helper.DebugModule,
			fx.Populate(&g),
			fx.Populate(&l),
			pkg.FakeModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)

	},
	)
}
