package pkg

import (
	"go.uber.org/fx"
)

// FakeFlavorDevModule -> Fake DI module for fx
var FakeFlavorDevModule = fx.Options(fx.Provide(FakeFlavorDevConstructorForDi))

// FakeFlavorReleaseModule -> Fake DI module for fx
var FakeFlavorReleaseModule = fx.Options(fx.Provide(FakeFlavorReleaseConstructorForDi))

// Flavor -> Data type for {flavor} that needed by LogHandler
type Flavor struct{ F string }

// FakeFlavorDevConstructorForDi -> Creating Fake tag needed by DI
func FakeFlavorDevConstructorForDi() *Flavor {
	return &Flavor{F: "dev"}
}

// FakeFlavorReleaseConstructorForDi -> Creating Fake tag needed by DI
func FakeFlavorReleaseConstructorForDi() *Flavor {
	return &Flavor{F: "release"}
}
