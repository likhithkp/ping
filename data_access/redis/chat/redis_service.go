package chat

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type ChatRedisService struct {
	client *redis.Client
}

func NewUrlRedisService(client *redis.Client) *ChatRedisService {
	return &ChatRedisService{client: client}
}

func (s *ChatRedisService) SetMessage(ctx context.Context, userId, longURL string) error {
	return s.client.HSet(ctx, "message:"+userId, longURL, 24*time.Hour).Err()
}

func (s *ChatRedisService) GetMessage(ctx context.Context, userId string) (map[string]string, error) {
	msg, err := s.client.HGetAll(ctx, "message:"+userId).Result()
	if err == redis.Nil {
		return nil, nil
	}
	return msg, err
}

func (s *ChatRedisService) DeleteMessage(ctx context.Context, userId string) error {
	_, err := s.client.HDel(ctx, "message:"+userId).Result()
	if err == redis.Nil {
		return nil
	}
	return err
}
