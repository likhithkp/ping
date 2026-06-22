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

func (repository *ChatRepository) SetMessage(ctx context.Context, userId, messageId, message string) error {
	err := repository.chatRedisService.SetMessage(ctx, userId, messageId, message)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ChatRepository) UpdateMessage(ctx context.Context, userId, messageId, message string) error {
	err := repository.chatRedisService.UpdateMessage(ctx, userId, messageId, message)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ChatRepository) GetMessage(ctx context.Context, userId string) ([]*chat.MessageEntity, error) {
	msg, err := repository.chatRedisService.GetMessage(ctx, userId)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (repository *ChatRepository) DeleteMessage(ctx context.Context, userId, messageId string) error {
	err := repository.chatRedisService.DeleteMessage(ctx, userId, messageId)
	if err != nil {
		return nil
	}

	return nil
}
