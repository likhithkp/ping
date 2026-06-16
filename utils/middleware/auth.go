package middleware

import (
	"strings"

	"github.com/likhithkp/ping/utils/jwt"
	"github.com/likhithkp/ping/utils/other"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtManager *jwt.GenerateJwtTokenManager, utils *other.Utils) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.Response(c, false, fiber.StatusUnauthorized, "Missing auth token", nil)
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtManager.ValidateJWT(tokenStr)
		if err != nil {
			return utils.Response(c, false, fiber.StatusUnauthorized, "Invalid or expired token", nil)
		}

		c.Locals("userID", claims["sub"])
		c.Locals("email", claims["email"])
		c.Locals("phonenumber", claims["phonenumber"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}

func NewAuthMiddleware(jwtManager *jwt.GenerateJwtTokenManager, utils *other.Utils) fiber.Handler {
	return AuthMiddleware(jwtManager, utils)
}
