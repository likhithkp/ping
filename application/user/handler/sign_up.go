package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/user/convertor"
	"github.com/likhithkp/ping/application/user/dto"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type SignUpHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	userRepository *user.UserRepository
}

func NewSignUpHandler(
	utils *other.Utils,
	logger *zap.Logger,
	userRepository *user.UserRepository,
) *SignUpHandler {
	return &SignUpHandler{
		utils:          utils,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (signUpHandler *SignUpHandler) SignUp(c *fiber.Ctx) error {
	newUser := new(dto.SignUpRequest)

	err := c.BodyParser(newUser)
	if err != nil {
		signUpHandler.logger.Error("failed to parse body", zap.Error(err))
		return signUpHandler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing signup body", nil)
	}

	if newUser.FirstName == "" ||
		newUser.LastName == "" ||
		newUser.UserName == "" ||
		newUser.Bio == "" ||
		newUser.DateOfBirth == "" ||
		newUser.Password == "" ||
		newUser.PhoneNumber == "" ||
		newUser.Email == "" {
		return signUpHandler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}

	userDomain, err := convertor.ConvertSignUpDtoToDomain(newUser)
	if err != nil {
		signUpHandler.logger.Error("failed to convert dto to user domain", zap.Error(err))
		return signUpHandler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		signUpHandler.logger.Error("failed to hash password", zap.Error(err))
		return signUpHandler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	userDomain.Password = string(hashedPassword)
	err = signUpHandler.userRepository.UpsertUser(c.Context(), userDomain)
	if err != nil {
		signUpHandler.logger.Error("failed to upsert user", zap.Error(err))
		return signUpHandler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	return signUpHandler.utils.Response(c, true, http.StatusCreated, "Sign up successful", nil)
}
