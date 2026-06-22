package handler

import (
	"encoding/json"

	"github.com/likhithkp/ping/application/chat/dto"
	"github.com/likhithkp/ping/data_access/repository/chat"
	connectionmap "github.com/likhithkp/ping/utils/connection_map"
	"github.com/likhithkp/ping/utils/ctx"
)

type AckMessageHandler struct {
	chatRepository *chat.ChatRepository
}

func NewAckMessageHandler(
	chatRepository *chat.ChatRepository,
) *AckMessageHandler {
	return &AckMessageHandler{
		chatRepository: chatRepository,
	}
}

func (handler *AckMessageHandler) AckMessage(userId string, message dto.Message) error {
	_, online := connectionmap.Connections[message.SenderId]
	if online {
		err := handler.chatRepository.DeleteMessage(ctx.Background, userId, message.Id)
		if err != nil {
			return err
		}
	} else {
		msgBytes, err := json.Marshal(message)
		if err != nil {
			return err
		}

		err = handler.chatRepository.SetMessage(ctx.Background, message.SenderId, message.Id, string(msgBytes))
		if err != nil {
			return err
		}
	}

	return nil
}
