package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/convertor"
	"github.com/likhithkp/ping/application/auth/dto"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type CreateChannelHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	storage        *storage.Uploader
	userRepository *user.UserRepository
}

func NewCreateChannelHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
) *CreateChannelHandler {
	return &CreateChannelHandler{
		utils:          utils,
		logger:         logger,
		storage:        storage,
		userRepository: userRepository,
	}
}

func (handler *CreateChannelHandler) CreateChannel(c *fiber.Ctx) error {
	newUser := new(dto.SignUpRequest)

	err := c.BodyParser(newUser)
	if err != nil {
		handler.logger.Error("failed to parse body", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing signup body", nil)
	}

	if newUser.FirstName == "" ||
		newUser.LastName == "" ||
		newUser.UserName == "" ||
		newUser.Bio == "" ||
		newUser.DateOfBirth == "" ||
		newUser.Password == "" ||
		newUser.PhoneNumber == "" ||
		newUser.Email == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}

	userDomain, err := convertor.ConvertSignUpDtoToDomain(newUser)
	if err != nil {
		handler.logger.Error("failed to convert dto to user domain", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	err = handler.userRepository.InsertUser(c.Context(), userDomain)
	if err != nil {
		handler.logger.Error("failed to upsert user", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	return handler.utils.Response(c, true, http.StatusCreated, "Sign up successful", nil)
}
