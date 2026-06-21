package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/likhithkp/ping/data_access/repository/chat"
	"github.com/likhithkp/ping/utils/config"
	"github.com/likhithkp/ping/utils/ctx"
)

type FetchOfflineMesagesHandler struct {
	chatRepository       *chat.ChatRepository
	sendWithRetryHandler *SendWithRetryHandler
}

func NewFetchOfflineMesagesHandler(
	chatRepository *chat.ChatRepository,
	sendWithRetryHandler *SendWithRetryHandler) *FetchOfflineMesagesHandler {
	return &FetchOfflineMesagesHandler{
		chatRepository:       chatRepository,
		sendWithRetryHandler: sendWithRetryHandler,
	}
}

func (handler *FetchOfflineMesagesHandler) FetchOfflineMesages(userId string, env *config.Env, conn *websocket.Conn) error {
	offlineMsg, err := handler.chatRepository.GetMessage(ctx.Background, userId)
	if err != nil {
		return err
	}

	offlineMsgBytes, err := json.Marshal(offlineMsg)
	if err != nil {
		return err
	}

	if len(offlineMsg) != 0 {
		err := conn.WriteMessage(websocket.TextMessage, offlineMsgBytes)
		if err != nil {
			retryCount, err := strconv.Atoi(env.RetryCount)
			if err != nil {
				return err
			}

			err = handler.sendWithRetryHandler.SendWithRetry(websocket.TextMessage, offlineMsgBytes, conn, retryCount)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
