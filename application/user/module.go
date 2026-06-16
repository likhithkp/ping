package user

import (
	"github.com/likhithkp/ping/application/user/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-user",
	fx.Provide(
		handler.NewSignUpHandler,
		handler.NewSignInHandler,
		NewUserController,
	),
	fx.Invoke(RegisterUserRoutes),
)
