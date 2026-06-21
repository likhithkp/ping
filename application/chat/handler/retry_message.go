package handler

import (
	"github.com/gofiber/contrib/websocket"
)

type SendWithRetryHandler struct{}

func NewSendWithRetryHandler() *SendWithRetryHandler {
	return &SendWithRetryHandler{}
}

func (handler *SendWithRetryHandler) SendWithRetry(mt int, msg []byte, conn *websocket.Conn, retryCount int) error {
	for i := 0; i < retryCount; i++ {
		err := conn.WriteMessage(mt, msg)
		retryCount--
		if err != nil {
			return err
		}
	}

	return nil
}
