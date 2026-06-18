package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/likhithkp/ping/application/chat/dto"
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/data_access/repository/chat"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/domain"
	"github.com/likhithkp/ping/utils/config"
	connectionmap "github.com/likhithkp/ping/utils/connection_map"
	_const "github.com/likhithkp/ping/utils/const"
	"github.com/likhithkp/ping/utils/ctx"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type WsConnectHandler struct {
	utils             *other.Utils
	env               *config.Env
	logger            *zap.Logger
	storage           *storage.Uploader
	userRepository    *user.UserRepository
	channelRepository *channel.ChannelRepository
	chatRepository    *chat.ChatRepository
}

func NewWsConnectHandler(
	utils *other.Utils,
	env *config.Env,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
	channelRepository *channel.ChannelRepository,
	chatRepository *chat.ChatRepository,
) *WsConnectHandler {
	return &WsConnectHandler{
		utils:             utils,
		env:               env,
		logger:            logger,
		storage:           storage,
		userRepository:    userRepository,
		channelRepository: channelRepository,
		chatRepository:    chatRepository,
	}
}

func (handler *WsConnectHandler) Ws(c *websocket.Conn) error {
	userId := c.Params("id")
	if userId == "" {
		return websocket.ErrBadHandshake
	}

	var (
		mt   int
		msg  []byte
		err  error
		conn *websocket.Conn
		ok   bool
	)

	connectionmap.Add(userId, c)
	defer connectionmap.Remove(userId)

	offlineMsgBytes, err := fetchOfflineMesages(userId, handler.chatRepository, handler.env, c)
	if err != nil {
		handler.logger.Error("failed to fetch offline messages", zap.Error(err))
	}

	for {
		//Receive
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

		fmt.Printf("%s", msg)

		var message dto.Message
		var channelDomain *domain.ChannelDomain

		err := json.Unmarshal(msg, &message)
		if err != nil {
			return websocket.ErrBadHandshake
		}

		switch message.Type {
		case _const.MESSAGE:
			channelDomain, err = handler.channelRepository.GetChannelById(context.Background(), message.ChannelId)
			if err != nil {
				handler.logger.Error("error while validating channel", zap.Error(err))
				return err
			}
			if channelDomain == nil {
				handler.logger.Error("channel not found", zap.Error(err))
				return websocket.ErrBadHandshake
			}

			var receiver string
			for _, user := range channelDomain.Users {
				if user.UserId != userId {
					receiver = user.UserId
					break
				}
			}

			conn, ok = connectionmap.Connections[receiver]
			ackMessage(mt, msg, conn)
		case _const.ACK:
			fmt.Printf("acked: %s", message.Id)
			err := handler.chatRepository.DeleteMessage(ctx.Background, message.SenderId, message.Id)
			if err != nil {
				handler.logger.Error("error while deleting the message from redis", zap.Error(err))
			}
		}

		//Send
		if ok {
			if message.Id == "" {
				message.Id = primitive.NewObjectID().Hex()
			}

			msg, err = json.Marshal(message)
			if err != nil {
				return websocket.ErrBadHandshake
			}

			if message.Type == _const.MESSAGE {
				err = handler.chatRepository.SetMessage(ctx.Background, userId, message.Id, string(msg))
				if err != nil {
					handler.logger.Error("error while storing message in redis", zap.Error(err))
				}
			}

			err = conn.WriteMessage(mt, msg)
			if err != nil {
				var retryCount int = 5
				retryCount, err := strconv.Atoi(handler.env.RetryCount)
				if err != nil {
					handler.logger.Error("failed to convert retry count to int", zap.Error(err))
				}

				err = sendWithRetry(websocket.TextMessage, offlineMsgBytes, c, retryCount)
				if err != nil {
					handler.logger.Error("failed to send offline mesages", zap.Error(err))
				}
			} else {
				go handler.chatRepository.DeleteMessage(ctx.Background, userId, message.Id)
			}
		} else {
			handler.chatRepository.SetMessage(ctx.Background, userId, message.Id, string(msg))
		}

	}

	return nil
}

func fetchOfflineMesages(userId string, chatRepository *chat.ChatRepository, env *config.Env, conn *websocket.Conn) ([]byte, error) {
	offlineMsg, err := chatRepository.GetMessage(ctx.Background, userId)
	if err != nil {
		return nil, err
	}

	offlineMsgBytes, err := json.Marshal(offlineMsg)
	if err != nil {
		return nil, err
	}

	if len(offlineMsg) != 0 {
		err := conn.WriteMessage(websocket.TextMessage, offlineMsgBytes)
		if err != nil {
			var retryCount int = 5
			retryCount, err := strconv.Atoi(env.RetryCount)
			if err != nil {
				return nil, err
			}

			err = sendWithRetry(websocket.TextMessage, offlineMsgBytes, conn, retryCount)
			if err != nil {
				return nil, err
			}
		} else {
			// go chatRepository.DeleteMessage(ctx.Background, userId, )
		}
	}

	return offlineMsgBytes, nil
}

func sendWithRetry(mt int, msg []byte, conn *websocket.Conn, retryCount int) error {
	for i := 0; i < retryCount; i++ {
		err := conn.WriteMessage(mt, msg)
		retryCount--
		if err != nil {
			return err
		}
	}

	return nil
}

func ackMessage(mt int, msg []byte, conn *websocket.Conn) error {
	var message *dto.Message
	err := json.Unmarshal(msg, message)
	if err != nil {
		return err
	}
	message.Type = _const.ACK

	ackMessage, err := json.Marshal(message)

	err = conn.WriteMessage(mt, ackMessage)
	if err != nil {
		return err
	}

	return nil
}
