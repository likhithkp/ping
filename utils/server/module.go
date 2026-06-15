package server

import "go.uber.org/fx"

var Module = fx.Module("utils-server",
	fx.Provide(NewFiberApp),
	fx.Invoke(RunHttpServer),
)
