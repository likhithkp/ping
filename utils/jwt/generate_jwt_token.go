package jwt

import (
	"time"

	"github.com/likhithkp/ping/utils/config"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v5"
)

type GenerateJwtTokenManager struct {
	config *config.Env
}

func NewGenerateJwtTokenManager(config *config.Env) *GenerateJwtTokenManager {
	return &GenerateJwtTokenManager{
		config: config,
	}
}

func (jwtManager *GenerateJwtTokenManager) GenerateJWT(id, email, phoneNumber string) (string, error) {
	month := (time.Hour * 24) * 30

	claims := jtoken.MapClaims{
		"sub":         id,
		"email":       email,
		"phonenumber": phoneNumber,
		"exp":         time.Now().UTC().Add(month * 6).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtManager.config.JwtSecret))
}

func (jwtManager *GenerateJwtTokenManager) ValidateJWT(tokenStr string) (jtoken.MapClaims, error) {
	token, err := jtoken.Parse(tokenStr, func(token *jtoken.Token) (interface{}, error) {
		if _, ok := token.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(jwtManager.config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jtoken.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
}
