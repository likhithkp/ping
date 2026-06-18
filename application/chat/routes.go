package chat

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterChatRoutes(app *fiber.App, controller *Controller) {
	app.Get("/ws/:id", controller.WebSocketUpgrade())
}
