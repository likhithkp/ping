package handler

import (
	"context"
	"encoding/json"
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
		err           error
		conn          *websocket.Conn
		online        bool
		message       dto.Message
		channelDomain *domain.ChannelDomain
		msgBytes      []byte
	)

	connectionmap.Add(userId, c)
	defer connectionmap.Remove(userId)
	defer c.Close()

	err = fetchOfflineMesages(userId, handler.chatRepository, handler.env, c)
	if err != nil {
		handler.logger.Error("failed to fetch offline messages", zap.Error(err))
	}

	for {
		//Receive
		_, msg, err := readMessage(c)
		if err != nil {
			handler.logger.Error("error while receiving message", zap.Error(err))
			return err
		}

		err = json.Unmarshal(msg, &message)
		if err != nil {
			return websocket.ErrBadHandshake
		}

		channelDomain, err = handler.channelRepository.GetChannelById(context.Background(), message.ChannelId)
		if err != nil {
			handler.logger.Error("error while validating channel", zap.Error(err))
			return err
		}
		if channelDomain == nil {
			handler.logger.Error("channel not found", zap.Error(err))
			return websocket.ErrBadHandshake
		}

		var sender string
		for _, user := range channelDomain.Users {
			if user.UserId != userId {
				sender = user.UserId
				break
			}
		}

		conn, online = connectionmap.Connections[sender]

		if message.Type == _const.ACK {
			err = ackMessage(message, handler.chatRepository)
			if err != nil {
				handler.logger.Error("error while acknowledging the message", zap.Error(err))
			}
		}

		//Send
		message.Id = primitive.NewObjectID().Hex()
		if online {
			msgBytes, err = json.Marshal(message)
			if err != nil {
				return websocket.ErrBadHandshake
			}

			if message.Type != _const.ACK {
				err = handler.chatRepository.SetMessage(ctx.Background, userId, message.Id, string(msgBytes))
				if err != nil {
					handler.logger.Error("error while storing message in redis", zap.Error(err))
				}
			}

			retryCount, err := strconv.Atoi(handler.env.RetryCount)
			if err != nil {
				handler.logger.Error("failed to convert retry count to int", zap.Error(err))
			}

			err = sendMessage(conn, websocket.TextMessage, msgBytes, retryCount)
			if err != nil {
				handler.logger.Error("error while sending message", zap.Error(err))
			}
		} else {
			//offline persistence
			handler.chatRepository.SetMessage(ctx.Background, userId, message.Id, string(msgBytes))
		}
	}
}

func readMessage(conn *websocket.Conn) (int, []byte, error) {
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

func sendMessage(conn *websocket.Conn, mt int, msg []byte, retryCount int) error {
	err := conn.WriteMessage(mt, msg)
	if err != nil {
		err = sendWithRetry(websocket.TextMessage, msg, conn, retryCount)
		if err != nil {
			return err
		}
	}

	return nil
}

func fetchOfflineMesages(userId string, chatRepository *chat.ChatRepository, env *config.Env, conn *websocket.Conn) error {
	offlineMsg, err := chatRepository.GetMessage(ctx.Background, userId)
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

			err = sendWithRetry(websocket.TextMessage, offlineMsgBytes, conn, retryCount)
			if err != nil {
				return err
			}
		}

		// var message dto.Message
		// for _, msg := range offlineMsg {
		// 	go ackMessage(msg, chatRepository)
		// }
	}

	return nil
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

func ackMessage(message dto.Message, chatRepository *chat.ChatRepository) error {
	_, online := connectionmap.Connections[message.SenderId]
	if online {
		err := chatRepository.DeleteMessage(ctx.Background, message.SenderId, message.Id)
		if err != nil {
			return err
		}
	} else {
		msgBytes, err := json.Marshal(message)
		if err != nil {
			return err
		}

		err = chatRepository.SetMessage(ctx.Background, message.SenderId, message.Id, string(msgBytes))
		if err != nil {
			return err
		}
	}

	return nil
}
