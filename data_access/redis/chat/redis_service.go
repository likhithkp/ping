package chat

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ChatRedisService struct {
	client *redis.Client
}

func NewChatRedisService(client *redis.Client) *ChatRedisService {
	return &ChatRedisService{client: client}
}

func (s *ChatRedisService) SetMessage(ctx context.Context, userId, messageId, message string) error {
	key := "message:" + userId
	err := s.client.HSet(ctx, key, messageId, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *ChatRedisService) UpdateMessage(ctx context.Context, userId, messageId string, message string) error {
	key := "message:" + userId
	err := s.client.HSet(ctx, key, messageId, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *ChatRedisService) GetMessage(ctx context.Context, userId string) (map[string]string, error) {
	msg, err := s.client.HGetAll(ctx, "message:"+userId).Result()
	if err == redis.Nil {
		return nil, nil
	}
	return msg, err
}

func (s *ChatRedisService) DeleteMessage(ctx context.Context, userId, messageId string) error {
	fmt.Println("Rached redis delete the acked message")
	fmt.Println("userID", userId)
	fmt.Println("messageId", messageId)
	_, err := s.client.HDel(ctx, "message:"+userId, messageId).Result()
	if err == redis.Nil {
		return nil
	}
	return err
}
