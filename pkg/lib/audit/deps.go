package audit

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	NewLogger,
	wire.Struct(new(Sink), "*"),
	wire.Struct(new(ReadStore), "*"),
	wire.Struct(new(WriteStore), "*"),
	wire.Struct(new(Query), "*"),
)
