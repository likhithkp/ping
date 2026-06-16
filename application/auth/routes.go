package auth

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(app *fiber.App, controller *UserController) {
	appGroup := app.Group("api/v1/auth")

	appGroup.Post("/login", controller.signInHandler.SignIn)
	appGroup.Post("/forgot-password", controller.forgotPasswordHandler.ForgotPassword)
	appGroup.Post("/verify-otp", controller.verifyOtpHandler.VerifyOtp)
	appGroup.Post("/reset-password", controller.resetPasswordHandler.ResetPassword)

	appGroup.Post("", controller.signUpHandler.SignUp)
}
