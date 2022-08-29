package helper_test

import (
	"fmt"
	"testing"

	"github.com/kaankoken/helper/pkg"
	"github.com/kaankoken/helper/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestReleaseLogger(t *testing.T) {
	logger := helper.InitializeLogger()
	tag := pkg.FakeConstructorForDi()
	l := helper.InitializeLoggerPtr(logger, tag)

	t.Parallel()

	t.Run("r-logger-error=no-error", func(t *testing.T) {
		t.Parallel()
		err := l.Error(nil)

		assert.Nil(t, err)
	})

	t.Run("r-logger-error=only-error", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		err := l.Error(fmt.Errorf(msg))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("r-logger-error=no-error&empty-callback", func(t *testing.T) {
		t.Parallel()

		f := func() {}
		err := l.ErrorWithCallback(nil, f)

		assert.Nil(t, err)
	})

	t.Run("r-logger-error=error&empty-callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		f := func() {}

		err := l.ErrorWithCallback(fmt.Errorf(msg), f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("r-logger-error=error&callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		f := func() {
			msg := fmt.Sprintf("%s %d", "Testing callback", 1)

			assert.Nil(t, nil)
			assert.NotNil(t, msg)
		}

		err := l.ErrorWithCallback(fmt.Errorf(msg), f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("r-logger-info=success", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		_msg := l.Info(msg)

		assert.Contains(t, _msg, msg)
	})

	t.Run("r-logger-info=success&empty-callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		f := func() {}

		_msg := l.InfoWithCallback(msg, f)

		assert.Contains(t, _msg, msg)
	})

	t.Run("r-logger-info=success&callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		f := func() {
			msg := fmt.Sprintf("%s %d", "Testing callback", 1)
			assert.Nil(t, nil)
			assert.NotNil(t, msg)
		}

		_msg := l.InfoWithCallback(msg, f)

		assert.Contains(t, _msg, msg)
	})
}

func TestReleaseLoggerWithFx(t *testing.T) {
	t.Parallel()

	t.Run("r-logger=injection-test", func(t *testing.T) {
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			helper.ReleaseModule,
			fx.Populate(&g),
			pkg.FakeModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("r-logger=injection-test-with-functions", func(t *testing.T) {
		var g fx.DotGraph
		var l *helper.RLogger

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			helper.ReleaseModule,
			fx.Populate(&g),
			fx.Populate(&l),
			pkg.FakeModule,
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)

		t.Run("r-logger-error=no-error", func(t *testing.T) {
			t.Parallel()
			err := l.Error(nil)

			assert.Nil(t, err)
		})

		t.Run("r-logger-error=only-error", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			err := l.Error(fmt.Errorf(msg))

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("r-logger-error=no-error&empty-callback", func(t *testing.T) {
			t.Parallel()

			f := func() {}
			err := l.ErrorWithCallback(nil, f)

			assert.Nil(t, err)
		})

		t.Run("r-logger-error=error&empty-callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			f := func() {}

			err := l.ErrorWithCallback(fmt.Errorf(msg), f)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("r-logger-error=error&callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			f := func() {
				msg := fmt.Sprintf("%s %d", "Testing callback", 1)

				assert.Nil(t, nil)
				assert.NotNil(t, msg)
			}

			err := l.ErrorWithCallback(fmt.Errorf(msg), f)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("r-logger-info=success", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			_msg := l.Info(msg)

			assert.Contains(t, _msg, msg)
		})

		t.Run("r-logger-info=success&empty-callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			f := func() {}

			_msg := l.InfoWithCallback(msg, f)

			assert.Contains(t, _msg, msg)
		})

		t.Run("r-logger-info=success&callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			f := func() {
				msg := fmt.Sprintf("%s %d", "Testing callback", 1)
				assert.Nil(t, nil)
				assert.NotNil(t, msg)
			}

			_msg := l.InfoWithCallback(msg, f)

			assert.Contains(t, _msg, msg)
		})
	})
}
