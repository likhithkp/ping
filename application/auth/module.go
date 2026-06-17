package auth

import (
	"github.com/likhithkp/ping/application/auth/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-auth",
	fx.Provide(
		handler.NewSignUpHandler,
		handler.NewSignInHandler,
		handler.NewForgotPasswordHandler,
		handler.NewVerifyOtpHandler,
		handler.NewResetPasswordHandler,
		NewController,
	),
	fx.Invoke(RegisterUserRoutes),
)
