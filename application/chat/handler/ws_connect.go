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
	utils                      *other.Utils
	env                        *config.Env
	logger                     *zap.Logger
	storage                    *storage.Uploader
	userRepository             *user.UserRepository
	channelRepository          *channel.ChannelRepository
	chatRepository             *chat.ChatRepository
	readMessageHandler         *ReadMessageHandler
	writeMessageHandler        *WriteMessageHandler
	fetchOfflineMesagesHandler *FetchOfflineMesagesHandler
	sendPongHandler            *SendPongHandler
	sendWithRetryHandler       *SendWithRetryHandler
	ackMessageHandler          *AckMessageHandler
}

func NewWsConnectHandler(
	utils *other.Utils,
	env *config.Env,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
	channelRepository *channel.ChannelRepository,
	chatRepository *chat.ChatRepository,
	readMessageHandler *ReadMessageHandler,
	writeMessageHandler *WriteMessageHandler,
	fetchOfflineMesagesHandler *FetchOfflineMesagesHandler,
	sendPongHandler *SendPongHandler,
	sendWithRetryHandler *SendWithRetryHandler,
	ackMessageHandler *AckMessageHandler,
) *WsConnectHandler {
	return &WsConnectHandler{
		utils:                      utils,
		env:                        env,
		logger:                     logger,
		storage:                    storage,
		userRepository:             userRepository,
		channelRepository:          channelRepository,
		chatRepository:             chatRepository,
		readMessageHandler:         readMessageHandler,
		writeMessageHandler:        writeMessageHandler,
		fetchOfflineMesagesHandler: fetchOfflineMesagesHandler,
		sendPongHandler:            sendPongHandler,
		sendWithRetryHandler:       sendWithRetryHandler,
		ackMessageHandler:          ackMessageHandler,
	}
}

func (handler *WsConnectHandler) Ws(c *websocket.Conn) error {
	userId := c.Params("id")
	if userId == "" {
		return websocket.ErrBadHandshake
	}

	var (
		err           error
		conn          *connectionmap.User
		online        bool
		message       dto.Message
		channelDomain *domain.ChannelDomain
		msgBytes      []byte
	)

	connectionmap.Add(userId, c)

	defer func() {
		handler.channelRepository.UpdateLastSeenForAllChannels(context.Background(), userId)
		connectionmap.Remove(userId)
		c.Close()
	}()

	err = handler.fetchOfflineMesagesHandler.FetchOfflineMesages(userId, handler.env, c)
	if err != nil {
		handler.logger.Error("failed to fetch offline messages", zap.Error(err))
	}

	for {
		//Receive
		_, msg, err := handler.readMessageHandler.ReadMessage(c)
		if err != nil {
			handler.logger.Error("error while receiving message", zap.Error(err))
			return err
		}

		fmt.Printf("msg: %s", msg)

		err = json.Unmarshal(msg, &message)
		if err != nil {
			return websocket.ErrBadHandshake
		}

		switch message.Type {
		case _const.ACK:
			err = handler.ackMessageHandler.AckMessage(message)
			if err != nil {
				handler.logger.Error("error while acknowledging the message", zap.Error(err))
			}
		case _const.PING:
			connectionmap.UpdateHeartbeat(userId)
			err = handler.sendPongHandler.SendPong(c)
			if err != nil {
				handler.logger.Error("error while acknowledging the ping", zap.Error(err))
			}
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

			var sender string
			for _, user := range channelDomain.Users {
				if user.UserId != userId {
					sender = user.UserId
					break
				}
			}
			conn, online = connectionmap.Connections[sender]
		}

		//Send
		message.Id = primitive.NewObjectID().Hex()
		message.SenderId = userId

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

			err = handler.writeMessageHandler.WriteMessage(conn.Ws, websocket.TextMessage, msgBytes, retryCount)
			if err != nil {
				handler.logger.Error("error while sending message", zap.Error(err))
			}
		} else {
			//offline persistence
			handler.chatRepository.SetMessage(ctx.Background, userId, message.Id, string(msgBytes))
		}
	}
}
