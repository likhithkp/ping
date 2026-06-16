package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/handler"
)

type UserController struct {
	signUpHandler         *handler.SignUpHandler
	signInHandler         *handler.SignInHandler
	forgotPasswordHandler *handler.ForgotPasswordHandler
	verifyOtpHandler      *handler.VerifyOtpHandler
	resetPasswordHandler  *handler.ResetPasswordHandler
}

func NewUserController(
	signUpHandler *handler.SignUpHandler,
	signInHandler *handler.SignInHandler,
	forgotPasswordHandler *handler.ForgotPasswordHandler,
	verifyOtpHandler *handler.VerifyOtpHandler,
	resetPasswordHandler *handler.ResetPasswordHandler,
) *UserController {
	return &UserController{
		signUpHandler:         signUpHandler,
		signInHandler:         signInHandler,
		forgotPasswordHandler: forgotPasswordHandler,
		verifyOtpHandler:      verifyOtpHandler,
		resetPasswordHandler:  resetPasswordHandler,
	}
}

func (userController *UserController) SignUp(c *fiber.Ctx) error {
	return userController.signUpHandler.SignUp(c)
}

func (userController *UserController) SignIn(c *fiber.Ctx) error {
	return userController.signInHandler.SignIn(c)
}

func (userController *UserController) ForgotPassword(c *fiber.Ctx) error {
	return userController.forgotPasswordHandler.ForgotPassword(c)
}

func (userController *UserController) VerifyOtp(c *fiber.Ctx) error {
	return userController.verifyOtpHandler.VerifyOtp(c)
}

func (userController *UserController) ResetPassword(c *fiber.Ctx) error {
	return userController.resetPasswordHandler.ResetPassword(c)
}
