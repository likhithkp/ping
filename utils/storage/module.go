package storage

import "go.uber.org/fx"

var Module = fx.Module("utils-storage",
	fx.Provide(
		NewUploader,
	),
)
