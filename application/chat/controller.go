package chat

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/chat/handler"
)

type Controller struct {
	wsHandler *handler.WsHandler
}

func NewController(
	wsHandler *handler.WsHandler,
) *Controller {
	return &Controller{
		wsHandler: wsHandler,
	}
}

func (controller *Controller) WebSocketUpgrade() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		controller.wsHandler.Ws(conn)
	})
}
