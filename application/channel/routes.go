package channel

import "github.com/gofiber/fiber/v2"

func RegisterChannelRoutes(app *fiber.App, controller *Controller, middleware fiber.Handler) {
	appGroup := app.Group("api/v1/channel")
	appGroup.Use(middleware)

	appGroup.Get("/list/:id", controller.getChannelListHandler.GetChannelList)
	appGroup.Get("/:id", controller.getChannelDetailsHandler.GetChannelDetails)
	appGroup.Delete("/:id", controller.deleteChannelHandler.DeleteChannel)

	appGroup.Post("", controller.createChannelHandler.CreateChannel)
}
