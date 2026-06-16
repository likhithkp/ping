package auth

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(app *fiber.App, controller *UserController) {
	appGroup := app.Group("api/v1/users")

	appGroup.Post("/login", controller.signInHandler.SignIn)

	appGroup.Post("", controller.signUpHandler.SignUp)
}
