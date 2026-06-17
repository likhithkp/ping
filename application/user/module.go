package user

import (
	"github.com/likhithkp/ping/application/user/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-user",
	fx.Provide(
		handler.NewGetProfileHandler,
		handler.NewEditProfileHandler,
		NewController,
	),
	fx.Invoke(RegisterUserRoutes),
)
