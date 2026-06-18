package chat

import (
	"context"

	"github.com/likhithkp/ping/data_access/redis/chat"
)

type ChatRepository struct {
	chatRedisService *chat.ChatRedisService
}

func NewChatRepository(
	chatRedisService *chat.ChatRedisService,
) *ChatRepository {
	return &ChatRepository{
		chatRedisService: chatRedisService,
	}
}

func (repository *ChatRepository) SetMessage(ctx context.Context, userId string, longUrl string) error {
	err := repository.chatRedisService.SetMessage(ctx, userId, longUrl)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ChatRepository) GetMessage(ctx context.Context, userId string) (map[string]string, error) {
	msg, err := repository.chatRedisService.GetMessage(ctx, userId)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (repository *ChatRepository) DeleteMessage(ctx context.Context, userId string) error {
	err := repository.chatRedisService.DeleteMessage(ctx, userId)
	if err != nil {
		return nil
	}

	return nil
}
