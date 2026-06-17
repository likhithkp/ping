package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/user/convertor"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type GetProfileHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	storage        *storage.Uploader
	userRepository *user.UserRepository
}

func NewGetProfileHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
) *GetProfileHandler {
	return &GetProfileHandler{
		utils:          utils,
		logger:         logger,
		storage:        storage,
		userRepository: userRepository,
	}
}

func (handler *GetProfileHandler) GetProfile(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "User id missing", nil)
	}

	userDomain, err := handler.userRepository.GetUserById(c.Context(), userId)
	if err != nil {
		handler.logger.Error("failed to upsert user", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}
	if userDomain == nil {
		return handler.utils.Response(c, false, http.StatusNotFound, "User not found", nil)
	}

	userDetails := convertor.ConvertDomainToGetDetails(userDomain)
	return handler.utils.Response(c, true, http.StatusCreated, "User details fetched successfully", userDetails)
}
