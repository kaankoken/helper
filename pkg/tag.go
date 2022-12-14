package pkg

import "go.uber.org/fx"

// FakeTagModule -> Fake DI module for fx
var FakeTagModule = fx.Options(fx.Provide(FakeConstructorForDi))

// Tag ->  Data type for {Tag} that needed by loggers
type Tag struct{ T string }

// FakeConstructorForDi -> Creating Fake tag needed by DI
func FakeConstructorForDi() *Tag {
	return &Tag{T: "Test Tag"}
}
