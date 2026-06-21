package handler

import "github.com/gofiber/contrib/websocket"

type WriteMessageHandler struct {
	sendWithRetryHandler *SendWithRetryHandler
}

func NewWriteMessageHandler(
	sendWithRetryHandler *SendWithRetryHandler,
) *WriteMessageHandler {
	return &WriteMessageHandler{
		sendWithRetryHandler: sendWithRetryHandler,
	}
}

func (handler *WriteMessageHandler) WriteMessage(conn *websocket.Conn, mt int, msg []byte, retryCount int) error {
	err := conn.WriteMessage(mt, msg)
	if err != nil {
		err = handler.sendWithRetryHandler.SendWithRetry(websocket.TextMessage, msg, conn, retryCount)
		if err != nil {
			return err
		}
	}

	return nil
}
