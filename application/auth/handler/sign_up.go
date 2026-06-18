package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/likhithkp/ping/application/auth/convertor"
	"github.com/likhithkp/ping/application/auth/dto"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/domain"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type SignUpHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	storage        *storage.Uploader
	userRepository *user.UserRepository
}

func NewSignUpHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
) *SignUpHandler {
	return &SignUpHandler{
		utils:          utils,
		logger:         logger,
		storage:        storage,
		userRepository: userRepository,
	}
}

func (handler *SignUpHandler) SignUp(c *fiber.Ctx) error {
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

	var existingDomain *domain.UserDomain
	existingDomain, err = handler.userRepository.GetUserByEmail(c.Context(), newUser.Email)
	if err != nil {
		handler.logger.Error("failed to check existing user by email", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}
	if existingDomain != nil {
		return handler.utils.Response(c, false, http.StatusConflict, "User with the email "+newUser.Email+" alredy exists", nil)
	}

	existingDomain, err = handler.userRepository.GetUserByPhoneNumber(c.Context(), newUser.PhoneNumber)
	if err != nil {
		handler.logger.Error("failed to check existing user by phonenumber", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}
	if existingDomain != nil {
		return handler.utils.Response(c, false, http.StatusConflict, "User with the number "+newUser.PhoneNumber+" alredy exists", nil)
	}

	userDomain, err := convertor.ConvertSignUpDtoToDomain(newUser)
	if err != nil {
		handler.logger.Error("failed to convert dto to user domain", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	if newUser.Image != "" {
		imageBytes, err := base64.StdEncoding.DecodeString(newUser.Image)
		if err != nil {
			handler.logger.Error("invalid base64 image format", zap.Error(err))
			return handler.utils.Response(c, false, http.StatusBadRequest, "Invalid image format. Please upload a valid photo.", nil)
		}

		cleanName := strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
				return r
			}
			return -1
		}, userDomain.UserName)
		if cleanName == "" {
			cleanName = uuid.NewString()
		}

		key := fmt.Sprintf("users/image/%s.jpg", cleanName)

		reader := bytes.NewReader(imageBytes)
		url, err := handler.storage.UploadFile(c.Context(), reader, key, "image/jpeg")
		if err != nil {
			handler.logger.Error("failed to upload image to S3", zap.Error(err))
			return handler.utils.Response(c, false, http.StatusInternalServerError, "Failed to upload image", nil)
		}

		userDomain.Image = url
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		handler.logger.Error("failed to hash password", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	userDomain.Password = string(hashedPassword)
	err = handler.userRepository.InsertUser(c.Context(), userDomain)
	if err != nil {
		handler.logger.Error("failed to upsert user", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	return handler.utils.Response(c, true, http.StatusCreated, "Sign up successful", nil)
}
