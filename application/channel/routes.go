package channel

import "github.com/gofiber/fiber/v2"

func RegisterChannelRoutes(app *fiber.App, controller *Controller) {
	appGroup := app.Group("api/v1/channel")

	appGroup.Get("/list/:id", controller.getChannelListHandler.GetChannelList)
	appGroup.Get("/:id", controller.getChannelDetailsHandler.GetChannelDetails)

	appGroup.Post("", controller.createChannelHandler.CreateChannel)
}
