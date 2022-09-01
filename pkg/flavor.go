package pkg

import (
	"go.uber.org/fx"
)

// FakeTagModule -> Fake DI module for fx
var FakeFlavorDevModule = fx.Options(fx.Provide(FakeFlavorDevConstructorForDi))

// FakeTagModule -> Fake DI module for fx
var FakeFlavorReleaseModule = fx.Options(fx.Provide(FakeFlavorReleaseConstructorForDi))

// Flavor -> Data type for {flavor} that needed by LogHandler
type Flavor struct{ F string }

// FakeFlavorConstructorForDi -> Creating Fake tag needed by DI
func FakeFlavorDevConstructorForDi() *Flavor {
	return &Flavor{F: "dev"}
}

// FakeFlavorConstructorForDi -> Creating Fake tag needed by DI
func FakeFlavorReleaseConstructorForDi() *Flavor {
	return &Flavor{F: "release"}
}
