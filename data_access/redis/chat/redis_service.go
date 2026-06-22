package chat

import (
	"context"
	"encoding/json"

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

func (s *ChatRedisService) GetMessage(ctx context.Context, userId string) ([]*MessageEntity, error) {
	msg, err := s.client.HGetAll(ctx, "message:"+userId).Result()
	if err == redis.Nil {
		return nil, nil
	}

	if len(msg) == 0 {
		return nil, nil
	}

	var messages []*MessageEntity
	for _, msgJSON := range msg {
		var message MessageEntity
		if err := json.Unmarshal([]byte(msgJSON), &message); err != nil {
			continue
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

func (s *ChatRedisService) DeleteMessage(ctx context.Context, userId, messageId string) error {
	_, err := s.client.HDel(ctx, "message:"+userId, messageId).Result()
	if err == redis.Nil {
		return nil
	}
	return err
}

func (s *ChatRedisService) DeleteAllMessage(ctx context.Context, userId string) error {
	err := s.client.Del(ctx, "message:"+userId).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}
