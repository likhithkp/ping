package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/channel/convertor"
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type GetChannelDetailsHandler struct {
	utils             *other.Utils
	logger            *zap.Logger
	storage           *storage.Uploader
	userRepository    *user.UserRepository
	channelRepository *channel.ChannelRepository
}

func NewGetChannelDetailsHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
	channelRepository *channel.ChannelRepository,
) *GetChannelDetailsHandler {
	return &GetChannelDetailsHandler{
		utils:             utils,
		logger:            logger,
		storage:           storage,
		userRepository:    userRepository,
		channelRepository: channelRepository,
	}
}

func (handler *GetChannelDetailsHandler) GetChannelDetails(c *fiber.Ctx) error {
	channelId := c.Params("id")
	if channelId == " " {
		return handler.utils.Response(c, true, http.StatusBadRequest, "Channel id missing", nil)
	}

	channelDomain, err := handler.channelRepository.GetChannelById(c.Context(), channelId)
	if err != nil {
		handler.logger.Error("failed to fetch channels", zap.Error(err))
		return handler.utils.Response(c, true, http.StatusInternalServerError, "Internal server error", nil)
	}
	if channelDomain == nil {
		return handler.utils.Response(c, true, http.StatusNotFound, "Channel not found", nil)
	}

	channelDetails := convertor.ConvertChannelDetailsToDto(channelDomain)
	return handler.utils.Response(c, true, http.StatusCreated, "Channels details fetched successfully", channelDetails)
}
