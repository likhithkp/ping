package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/likhithkp/ping/application/user/convertor"
	"github.com/likhithkp/ping/application/user/dto"
	"github.com/likhithkp/ping/data_access/repository/user"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/zap"
)

type EditProfileHandler struct {
	utils          *other.Utils
	logger         *zap.Logger
	storage        *storage.Uploader
	userRepository *user.UserRepository
}

func NewEditProfileHandler(
	utils *other.Utils,
	logger *zap.Logger,
	storage *storage.Uploader,
	userRepository *user.UserRepository,
) *EditProfileHandler {
	return &EditProfileHandler{
		utils:          utils,
		logger:         logger,
		storage:        storage,
		userRepository: userRepository,
	}
}

func (handler *EditProfileHandler) EditProfile(c *fiber.Ctx) error {
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

	editUser := new(dto.EditProfileRequest)
	err = c.BodyParser(editUser)
	if err != nil {
		handler.logger.Error("failed to parse body", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusUnprocessableEntity, "Error while parsing request body", nil)
	}

	if editUser.FirstName == "" ||
		editUser.LastName == "" ||
		editUser.UserName == "" ||
		editUser.Bio == "" ||
		editUser.DateOfBirth == "" ||
		editUser.PhoneNumber == "" ||
		editUser.Email == "" {
		return handler.utils.Response(c, false, http.StatusBadRequest, "Missing fields", nil)
	}

	updatedUserDomain, err := convertor.ConvertEditProfileDtoToDomain(userId, editUser)
	if err != nil {
		handler.logger.Error("failed to convert dto to user domain", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Internal server error", nil)
	}

	if editUser.NewImage != "" {
		imageBytes, err := base64.StdEncoding.DecodeString(editUser.NewImage)
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

		updatedUserDomain.Image = url
	} else {
		updatedUserDomain.Image = userDomain.Image
	}

	updatedUserDomain.Password = userDomain.Password
	err = handler.userRepository.UpdateUser(c.Context(), updatedUserDomain)
	if err != nil {
		handler.logger.Error("failed to upsert user", zap.Error(err))
		return handler.utils.Response(c, false, http.StatusInternalServerError, "Failed to update user", nil)
	}

	return handler.utils.Response(c, true, http.StatusCreated, "Profile updated successfully", nil)
}
