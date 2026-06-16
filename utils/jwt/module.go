package jwt

import "go.uber.org/fx"

var Module = fx.Module("utils-jwt",
	fx.Provide(
		NewGenerateJwtTokenManager,
		NewVerifyJwtTokenManager,
	),
)
