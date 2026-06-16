package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/dto"
	"github.com/likhithkp/ping/data_access/repository/otp"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/domain"
	"github.com/likhithkp/ping/utils/jwt"
	"github.com/likhithkp/ping/utils/mail"
	"github.com/likhithkp/ping/utils/other"
	"go.uber.org/zap"
)

type VerifyOtpHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	jwtManager     *jwt.GenerateJwtTokenManager
	mailer         *mail.Mailer
	userRepository *user.UserRepository
	otpRepository  *otp.OtpRepository
}

func NewVerifyOtpHandler(
	utils *other.Utils,
	logger *zap.Logger,
	jwtManager *jwt.GenerateJwtTokenManager,
	mailer *mail.Mailer,
	userRepository *user.UserRepository,
	otpRepository *otp.OtpRepository,
) *VerifyOtpHandler {
	return &VerifyOtpHandler{
		utils:          utils,
		logger:         logger,
		jwtManager:     jwtManager,
		mailer:         mailer,
		userRepository: userRepository,
		otpRepository:  otpRepository,
	}
}

func (handler *VerifyOtpHandler) VerifyOtp(c *fiber.Ctx) error {
	otpDto := new(dto.VerifyOtpRequest)

	err := c.BodyParser(otpDto)
	if err != nil {
		handler.logger.Error("failed to parse body", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing forgot password body", nil)
	}

	if otpDto.Email == "" || otpDto.Otp == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}

	var userDomain *domain.UserDomain
	userDomain, err = handler.userRepository.GetUserByEmail(c.Context(), otpDto.Email)
	if err != nil {
		handler.logger.Error("failed to get user by email", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	if userDomain == nil {
		return handler.utils.Response(c, false, http.StatusNotFound, "User with email "+otpDto.Email+" does not exist", nil)
	}

	otpDomain, err := handler.otpRepository.GetLatestOtpByEmail(c.Context(), otpDto.Email)
	if err != nil {
		handler.logger.Error("failed to fetch otp", zap.Error(err))
		return handler.utils.Response(c, false, fiber.StatusInternalServerError, "Internal server error", nil)
	}
	fmt.Println(otpDomain)
	if otpDomain == nil || otpDto.Otp != otpDomain.Otp {
		return handler.utils.Response(c, false, fiber.StatusUnauthorized, "Invalid OTP", nil)
	}
	if time.Since(otpDomain.CreateAt) > 10*time.Minute {
		return handler.utils.Response(c, false, fiber.StatusGone, "OTP expired", nil)
	}

	return handler.utils.Response(c, true, http.StatusOK, "OTP verified successfully", nil)
}
