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

type GetChannelListHandler struct {
	utils             *other.Utils
	logger            *zap.Logger
	storage           *storage.Uploader
	userRepository    *user.UserRepository
	channelRepository *channel.ChannelRepository
}

func NewGetChannelListHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
	channelRepository *channel.ChannelRepository,
) *GetChannelListHandler {
	return &GetChannelListHandler{
		utils:             utils,
		logger:            logger,
		storage:           storage,
		userRepository:    userRepository,
		channelRepository: channelRepository,
	}
}

func (handler *GetChannelListHandler) GetChannelList(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == " " {
		return handler.utils.Response(c, true, http.StatusBadRequest, "User id missing", nil)
	}

	channelDomain, err := handler.channelRepository.GetChannelsByUserId(c.Context(), userId)
	if err != nil {
		handler.logger.Error("failed to fetch channels", zap.Error(err))
		return handler.utils.Response(c, true, http.StatusInternalServerError, "Internal server error", nil)
	}
	if len(channelDomain) == 0 {
		return handler.utils.Response(c, true, http.StatusOK, "No channels found", nil)
	}

	channelList := convertor.ConvertChannelListToDto(channelDomain)
	return handler.utils.Response(c, true, http.StatusCreated, "Channels fetched successfully", channelList)
}
