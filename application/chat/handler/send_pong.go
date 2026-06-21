package handler

import (
	"encoding/json"

	"github.com/gofiber/contrib/websocket"
	"github.com/likhithkp/ping/application/chat/dto"
	_const "github.com/likhithkp/ping/utils/const"
)

type SendPongHandler struct{}

func NewSendPongHandler() *SendPongHandler {
	return &SendPongHandler{}
}

func (handler *SendPongHandler) SendPong(conn *websocket.Conn) error {
	msg := dto.Message{Type: _const.PONG}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		return err
	}

	return nil
}
