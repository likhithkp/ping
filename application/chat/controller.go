package chat

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/chat/handler"
)

type Controller struct {
	WsConnectHandler *handler.WsConnectHandler
}

func NewController(
	WsConnectHandler *handler.WsConnectHandler,
) *Controller {
	return &Controller{
		WsConnectHandler: WsConnectHandler,
	}
}

func (controller *Controller) WebSocketUpgrade() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		controller.WsConnectHandler.Ws(conn)
	})
}
