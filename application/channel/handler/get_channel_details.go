package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type GetChannelDetailsHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	storage        *storage.Uploader
	userRepository *user.UserRepository
}

func NewGetChannelDetailsHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
) *GetChannelDetailsHandler {
	return &GetChannelDetailsHandler{
		utils:          utils,
		logger:         logger,
		storage:        storage,
		userRepository: userRepository,
	}
}

func (handler *GetChannelDetailsHandler) GetChannelDetails(c *fiber.Ctx) error {
	return handler.utils.Response(c, true, http.StatusCreated, "Sign up successful", nil)
}
