package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/dto"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/domain"
	"github.com/likhithkp/ping/utils/other"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	userRepository *user.UserRepository
}

func NewResetPasswordHandler(
	utils *other.Utils,
	logger *zap.Logger,
	userRepository *user.UserRepository,
) *ResetPasswordHandler {
	return &ResetPasswordHandler{
		utils:          utils,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (handler *ResetPasswordHandler) ResetPassword(c *fiber.Ctx) error {
	req := new(dto.ResetPasswordRequest)

	err := c.BodyParser(req)
	if err != nil {
		handler.logger.Error("failed to parse body", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing signin body", nil)
	}

	if req.NewPassword == "" ||
		req.OldPassword == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}

	var userDomain *domain.UserDomain
	userDomain, err = handler.userRepository.GetUserByEmail(c.Context(), req.Email)
	if err != nil {
		handler.logger.Error("failed to get req by email", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}
	if userDomain == nil {
		return handler.utils.Response(c, false, http.StatusNotFound, "User with email "+req.Email+" does not exist", nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(req.OldPassword))
	if err != nil {
		handler.logger.Error("failed to compare old password", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Invalid old password", nil)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		handler.logger.Error("failed to generate new  password", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	userDomain.Password = string(hashedPassword)
	err = handler.userRepository.UpdateUser(c.Context(), userDomain)
	if err != nil {
		handler.logger.Error("failed to update password", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	return handler.utils.Response(c, true, http.StatusOK, "Password reset successful", nil)
}
