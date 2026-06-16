package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/auth/dto"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/domain"
	_const "github.com/likhithkp/ping/utils/const"
	"github.com/likhithkp/ping/utils/jwt"
	"github.com/likhithkp/ping/utils/other"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type SignInHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	userRepository *user.UserRepository
	jwtManager     *jwt.GenerateJwtTokenManager
}

func NewSignInHandler(
	utils *other.Utils,
	logger *zap.Logger,
	userRepository *user.UserRepository,
	jwtManager *jwt.GenerateJwtTokenManager,
) *SignInHandler {
	return &SignInHandler{
		utils:          utils,
		logger:         logger,
		userRepository: userRepository,
		jwtManager:     jwtManager,
	}
}

func (signInHandler *SignInHandler) SignIn(c *fiber.Ctx) error {
	user := new(dto.SignInRequest)

	err := c.BodyParser(user)
	if err != nil {
		signInHandler.logger.Error("failed to parse body", zap.Error(err))
		return signInHandler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing signin body", nil)
	}

	if user.Password == "" ||
		user.IdentifierType == "" {
		return signInHandler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}
	if user.IdentifierType == _const.EMAIL && user.Email == "" {
		return signInHandler.utils.Response(c, false, http.StatusBadRequest, "Missing email", nil)
	}
	if user.IdentifierType == _const.PHONE && user.PhoneNumber == "" {
		return signInHandler.utils.Response(c, false, http.StatusBadRequest, "Missing phonenumber", nil)
	}

	var userDomain *domain.UserDomain
	if user.IdentifierType == _const.EMAIL {
		userDomain, err = signInHandler.userRepository.GetUserByEmail(c.Context(), user.Email)
		if err != nil {
			signInHandler.logger.Error("failed to get user by email", zap.Error(err))
			return signInHandler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
		}

		if userDomain == nil {
			return signInHandler.utils.Response(c, false, http.StatusNotFound, "User with email "+user.Email+" does not exist", nil)
		}
	} else {
		userDomain, err = signInHandler.userRepository.GetUserByPhoneNumber(c.Context(), user.PhoneNumber)
		if err != nil {
			signInHandler.logger.Error("failed to get user by phonenumber", zap.Error(err))
			return signInHandler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
		}

		if userDomain == nil {
			return signInHandler.utils.Response(c, false, http.StatusNotFound, "User with number "+user.PhoneNumber+" does not exist", nil)
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(user.Password))
	if err != nil {
		signInHandler.logger.Error("failed to compare hash and password", zap.Error(err))
		return signInHandler.utils.Response(c, false, http.StatusUnauthorized, "Invalid password", nil)
	}

	token, err := signInHandler.jwtManager.GenerateJWT(userDomain.Id, userDomain.Email, userDomain.PhoneNumber)
	if err != nil {
		signInHandler.logger.Error("failed to generate jwt token", zap.Error(err))
		return signInHandler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	return signInHandler.utils.Response(c, true, http.StatusOK, "Sign in successful", fiber.Map{
		"token": token,
	})
}
