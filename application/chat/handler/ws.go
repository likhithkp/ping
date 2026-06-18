package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/likhithkp/ping/application/chat/dto"
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/data_access/repository/chat"
	"github.com/likhithkp/ping/data_access/repository/user"
	connectionmap "github.com/likhithkp/ping/utils/connection_map"
	"github.com/likhithkp/ping/utils/ctx"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type WsHandler struct {
	utils             *other.Utils
	logger            *zap.Logger
	storage           *storage.Uploader
	userRepository    *user.UserRepository
	channelRepository *channel.ChannelRepository
	chatRepository    *chat.ChatRepository
}

func NewWsHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
	channelRepository *channel.ChannelRepository,
	chatRepository *chat.ChatRepository,
) *WsHandler {
	return &WsHandler{
		utils:             utils,
		logger:            logger,
		storage:           storage,
		userRepository:    userRepository,
		channelRepository: channelRepository,
		chatRepository:    chatRepository,
	}
}

func (handler *WsHandler) Ws(c *websocket.Conn) error {
	userId := c.Params("id")
	if userId == "" {
		return websocket.ErrBadHandshake
	}

	connectionmap.Add(userId, c)
	defer connectionmap.Remove(userId)

	offlineMsg, err := handler.chatRepository.GetMessage(ctx.Background, userId)
	if err != nil {
		handler.logger.Error("failed to fetch offline messsages from redis", zap.Error(err))
	}

	offlineMsgBytes, err := json.Marshal(offlineMsg)
	if err != nil {
		handler.logger.Error("failed to convert offline messsages to bytes", zap.Error(err))
	}
	if len(offlineMsg) != 0 {
		err = c.WriteMessage(websocket.TextMessage, offlineMsgBytes)
		if err != nil {
			//TODO:try 5 times for delivery, after 5 retires if it fails drop it
			sendWithRetry()
			handler.logger.Error("failed to send offline mesages", zap.Error(err))
		} else {
			go handler.chatRepository.DeleteMessage(ctx.Background, userId)
		}
	}

	var (
		mt  int
		msg []byte
	)
	for {
		mt, msg, err = c.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				handler.logger.Error("websocket disconnected", zap.Error(err))
			} else if websocket.IsUnexpectedCloseError(err) {
				handler.logger.Error("websocket closed unexpectedly", zap.Error(err))
			} else {
				handler.logger.Error("websocket disconnected due to an internal server error", zap.Error(err))
			}
			break
		}

		var message dto.Message
		err := json.Unmarshal(msg, &message)
		if err != nil {
			return websocket.ErrBadHandshake
		}
		message.Id = primitive.NewObjectID().Hex()

		channelDomain, err := handler.channelRepository.GetChannelById(context.Background(), message.ChannelId)
		if err != nil {
			handler.logger.Error("error while validating channel", zap.Error(err))
			return err
		}
		if channelDomain == nil {
			handler.logger.Error("channel not found", zap.Error(err))
			return websocket.ErrBadHandshake
		}

		var recevier string
		for _, user := range channelDomain.Users {
			if user.UserId != userId {
				recevier = user.UserId
				break
			}
		}

		conn, ok := connectionmap.Connections[recevier]
		if ok {
			msg, err = json.Marshal(message)
			if err != nil {
				return websocket.ErrBadHandshake
			}

			handler.chatRepository.SetMessage(ctx.Background, userId, string(msg))

			//Store in redis first, after getting the ACK for that message delete it immediatety, if not retry 5 times and if it still faild drop it
			err = conn.WriteMessage(mt, msg)
			if err != nil {

			}
		} else {
			//TODO:Save in redis
			fmt.Printf("Receiver %s offline, msg saved in redis \n", recevier)
		}
	}

	return nil
}

func sendWithRetry() {}
