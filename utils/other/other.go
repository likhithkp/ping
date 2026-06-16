package other

import (
	"github.com/gofiber/fiber/v2"
)

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (o *Utils) Response(c *fiber.Ctx, success bool, status int, message string, data interface{}) error {
	return c.Status(status).JSON(Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}
