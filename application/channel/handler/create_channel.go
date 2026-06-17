package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/channel/convertor"
	"github.com/likhithkp/ping/application/channel/dto"
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type CreateChannelHandler struct {
	utils             *other.Utils
	logger            *zap.Logger
	storage           *storage.Uploader
	userRepository    *user.UserRepository
	channelRepository *channel.ChannelRepository
}

func NewCreateChannelHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
	channelRepository *channel.ChannelRepository,
) *CreateChannelHandler {
	return &CreateChannelHandler{
		utils:             utils,
		logger:            logger,
		storage:           storage,
		userRepository:    userRepository,
		channelRepository: channelRepository,
	}
}

func (handler *CreateChannelHandler) CreateChannel(c *fiber.Ctx) error {
	newChannel := new(dto.ChannelDto)

	err := c.BodyParser(newChannel)
	if err != nil {
		handler.logger.Error("failed to parse body", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing request body", nil)
	}
	if len(newChannel.Users) <= 1 {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Channel must have minimum of 2 members", nil)
	}

	userDomains, err := handler.userRepository.GetUsersByIDs(c.Context(), newChannel.Users)
	if err != nil {
		handler.logger.Error("failed to fetch users by ids", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	newChannelDomain := convertor.ConvertCreateChannelDtoToDomain(userDomains)
	err = handler.channelRepository.InsertChannel(c.Context(), newChannelDomain)
	if err != nil {
		handler.logger.Error("failed to upsert channel", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	return handler.utils.Response(c, true, http.StatusCreated, "Channel created successfully", nil)
}
