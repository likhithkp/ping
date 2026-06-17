package user

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(app *fiber.App, controller *Controller) {
	appGroup := app.Group("api/v1/user")

	appGroup.Get("/:id", controller.getProfileHandler.GetProfile)
	appGroup.Put("/:id", controller.editProfileHandler.EditProfile)
}
