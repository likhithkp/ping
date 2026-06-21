package handler

import "github.com/gofiber/contrib/websocket"

type ReadMessageHandler struct{}

func NewReadMessageHandler() *ReadMessageHandler {
	return &ReadMessageHandler{}
}

func (handler *ReadMessageHandler) ReadMessage(conn *websocket.Conn) (int, []byte, error) {
	mt, msg, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			return 0, nil, err
		} else if websocket.IsUnexpectedCloseError(err) {
			return 0, nil, err
		} else {
			return 0, nil, err
		}
	}

	return mt, msg, err
}
