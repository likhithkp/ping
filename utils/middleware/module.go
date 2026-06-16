package middleware

import "go.uber.org/fx"

var Module = fx.Module("utils-middleware",
	fx.Provide(
		NewAuthMiddleware,
	),
)
