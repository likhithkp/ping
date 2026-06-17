package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/handler"
)

type Controller struct {
	signUpHandler         *handler.SignUpHandler
	signInHandler         *handler.SignInHandler
	forgotPasswordHandler *handler.ForgotPasswordHandler
	verifyOtpHandler      *handler.VerifyOtpHandler
	resetPasswordHandler  *handler.ResetPasswordHandler
}

func NewController(
	signUpHandler *handler.SignUpHandler,
	signInHandler *handler.SignInHandler,
	forgotPasswordHandler *handler.ForgotPasswordHandler,
	verifyOtpHandler *handler.VerifyOtpHandler,
	resetPasswordHandler *handler.ResetPasswordHandler,
) *Controller {
	return &Controller{
		signUpHandler:         signUpHandler,
		signInHandler:         signInHandler,
		forgotPasswordHandler: forgotPasswordHandler,
		verifyOtpHandler:      verifyOtpHandler,
		resetPasswordHandler:  resetPasswordHandler,
	}
}

func (controller *Controller) SignUp(c *fiber.Ctx) error {
	return controller.signUpHandler.SignUp(c)
}

func (controller *Controller) SignIn(c *fiber.Ctx) error {
	return controller.signInHandler.SignIn(c)
}

func (controller *Controller) ForgotPassword(c *fiber.Ctx) error {
	return controller.forgotPasswordHandler.ForgotPassword(c)
}

func (controller *Controller) VerifyOtp(c *fiber.Ctx) error {
	return controller.verifyOtpHandler.VerifyOtp(c)
}

func (controller *Controller) ResetPassword(c *fiber.Ctx) error {
	return controller.resetPasswordHandler.ResetPassword(c)
}
