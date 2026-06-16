package auth

import (
	"github.com/likhithkp/ping/application/auth/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-user",
	fx.Provide(
		handler.NewSignUpHandler,
		handler.NewSignInHandler,
		handler.NewForgotPasswordHandler,
		handler.NewVerifyOtpHandler,
		handler.NewResetPasswordHandler,
		NewUserController,
	),
	fx.Invoke(RegisterUserRoutes),
)
