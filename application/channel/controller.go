package channel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/channel/handler"
)

type Controller struct {
	createChannelHandler     *handler.CreateChannelHandler
	getChannelListHandler    *handler.GetChannelListHandler
	getChannelDetailsHandler *handler.GetChannelDetailsHandler
}

func NewController(
	createChannelHandler *handler.CreateChannelHandler,
	getChannelListHandler *handler.GetChannelListHandler,
	getChannelDetailsHandler *handler.GetChannelDetailsHandler,
) *Controller {
	return &Controller{
		createChannelHandler:     createChannelHandler,
		getChannelListHandler:    getChannelListHandler,
		getChannelDetailsHandler: getChannelDetailsHandler,
	}
}

func (controller *Controller) CreateChannel(c *fiber.Ctx) error {
	return controller.createChannelHandler.CreateChannel(c)
}

func (controller *Controller) GetChannelList(c *fiber.Ctx) error {
	return controller.getChannelListHandler.GetChannelList(c)
}

func (controller *Controller) getChannelDetails(c *fiber.Ctx) error {
	return controller.getChannelDetailsHandler.GetChannelDetails(c)
}
