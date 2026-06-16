package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/convertor"
	"github.com/likhithkp/ping/application/auth/dto"
	"github.com/likhithkp/ping/data_access/repository/otp"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/domain"
	_const "github.com/likhithkp/ping/utils/const"
	"github.com/likhithkp/ping/utils/helper"
	"github.com/likhithkp/ping/utils/jwt"
	"github.com/likhithkp/ping/utils/mail"
	"github.com/likhithkp/ping/utils/other"
	"go.uber.org/zap"
)

type ForgotPasswordHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	jwtManager     *jwt.GenerateJwtTokenManager
	mailer         *mail.Mailer
	userRepository *user.UserRepository
	otpRepository  *otp.OtpRepository
}

func NewForgotPasswordHandler(
	utils *other.Utils,
	logger *zap.Logger,
	jwtManager *jwt.GenerateJwtTokenManager,
	mailer *mail.Mailer,
	userRepository *user.UserRepository,
	otpRepository *otp.OtpRepository,
) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		utils:          utils,
		logger:         logger,
		jwtManager:     jwtManager,
		mailer:         mailer,
		userRepository: userRepository,
		otpRepository:  otpRepository,
	}
}

func (handler *ForgotPasswordHandler) ForgotPassword(c *fiber.Ctx) error {
	user := new(dto.ForgotPasswordRequest)

	err := c.BodyParser(user)
	if err != nil {
		handler.logger.Error("failed to parse body", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing forgot password body", nil)
	}

	if user.IdentifierType == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}
	if user.IdentifierType == _const.EMAIL && user.Email == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing email", nil)
	}
	if user.IdentifierType == _const.PHONE && user.PhoneNumber == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing phonenumber", nil)
	}

	var userDomain *domain.UserDomain
	if user.IdentifierType == _const.EMAIL {
		userDomain, err = handler.userRepository.GetUserByEmail(c.Context(), user.Email)
		if err != nil {
			handler.logger.Error("failed to get user by email", zap.Error(err))
			return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
		}

		if userDomain == nil {
			return handler.utils.Response(c, false, http.StatusNotFound, "User with email "+user.Email+" does not exist", nil)
		}
	} else {
		userDomain, err = handler.userRepository.GetUserByPhoneNumber(c.Context(), user.PhoneNumber)
		if err != nil {
			handler.logger.Error("failed to get user by phonenumber", zap.Error(err))
			return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
		}

		if userDomain == nil {
			return handler.utils.Response(c, false, http.StatusNotFound, "User with number "+user.PhoneNumber+" does not exist", nil)
		}
	}

	otp := helper.GenerateOTP()
	optDomain := convertor.ConvertOtpReqToDomain(userDomain.Email, userDomain.Id, otp)
	err = handler.otpRepository.UpsertOtp(c.Context(), optDomain)
	if err != nil {
		handler.logger.Error("failed to upsert otp", zap.Error(err))
		return handler.utils.Response(c, false, fiber.StatusInternalServerError, "Internal server error", nil)
	}

	err = handler.mailer.SendEmail(c.Context(), userDomain.Email, otp)
	if err != nil {
		handler.logger.Error("failed to send otp", zap.Error(err))
		return handler.utils.Response(c, false, fiber.StatusBadGateway, "Failed to send OTP", nil)
	}

	return handler.utils.Response(c, true, http.StatusOK, "OTP sent successfully", nil)
}
