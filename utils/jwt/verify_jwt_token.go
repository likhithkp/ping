package jwt

import (
	"errors"

	"github.com/likhithkp/ping/utils/config"

	"github.com/golang-jwt/jwt/v5"
)

type VerifyJwtTokenManager struct {
	config *config.Env
}

func NewVerifyJwtTokenManager(config *config.Env) *VerifyJwtTokenManager {
	return &VerifyJwtTokenManager{
		config: config,
	}
}

func (verifyJwtTokenManager *VerifyJwtTokenManager) VerifyToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(verifyJwtTokenManager.config.JwtSecret), nil
	})
}
