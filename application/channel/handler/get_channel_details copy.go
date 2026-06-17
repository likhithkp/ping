package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type DeleteChannelHandler struct {
	utils             *other.Utils
	logger            *zap.Logger
	storage           *storage.Uploader
	channelRepository *channel.ChannelRepository
}

func NewDeleteChannelHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	channelRepository *channel.ChannelRepository,
) *DeleteChannelHandler {
	return &DeleteChannelHandler{
		utils:             utils,
		logger:            logger,
		storage:           storage,
		channelRepository: channelRepository,
	}
}

func (handler *DeleteChannelHandler) DeleteChannel(c *fiber.Ctx) error {
	channelId := c.Params("id")
	if channelId == " " {
		return handler.utils.Response(c, true, http.StatusBadRequest, "Channel id missing", nil)
	}

	err := handler.channelRepository.DeleteChannel(c.Context(), channelId)
	if err != nil {
		handler.logger.Error("failed to delete channel", zap.Error(err))
		return handler.utils.Response(c, true, http.StatusInternalServerError, "Internal server error", nil)
	}

	return handler.utils.Response(c, true, http.StatusOK, "Channels deleted  successfully", nil)
}
