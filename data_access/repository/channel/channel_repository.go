package channel

import (
	"context"

	channelMongoService "github.com/likhithkp/ping/data_access/mongo/channel"
	"github.com/likhithkp/ping/data_access/repository/channel/convertor"
	"github.com/likhithkp/ping/domain"
)

type ChannelRepository struct {
	channelMongoService *channelMongoService.ChannelMongoService
}

func NewChannelRepository(channelMongoService *channelMongoService.ChannelMongoService) *ChannelRepository {
	return &ChannelRepository{
		channelMongoService: channelMongoService,
	}
}

func (repository *ChannelRepository) UpdateChannel(ctx context.Context, channelDomain *domain.ChannelDomain) error {
	channelEntity, err := convertor.ConvertDomainToEntity(*channelDomain)
	if err != nil {
		return err
	}

	err = repository.channelMongoService.InsertChannel(ctx, channelEntity)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ChannelRepository) InsertChannel(ctx context.Context, channelDomain *domain.ChannelDomain) error {
	channelEntity, err := convertor.ConvertDomainToEntity(*channelDomain)
	if err != nil {
		return err
	}

	err = repository.channelMongoService.InsertChannel(ctx, channelEntity)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ChannelRepository) GetChannelById(ctx context.Context, id string) (channelDomain *domain.ChannelDomain, err error) {
	channelEntity, err := repository.channelMongoService.GetChannelById(ctx, id)
	if err != nil {
		return nil, err
	}
	if channelEntity == nil {
		return nil, nil
	}

	channelDomain, err = convertor.ConvertEntityToDomain(channelEntity)
	if err != nil {
		return nil, err
	}

	return channelDomain, nil
}

func (repository *ChannelRepository) GetChannelsByUserId(ctx context.Context, id string) (channelDomain []*domain.ChannelDomain, err error) {
	channelEntity, err := repository.channelMongoService.GetChannelsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	if channelEntity == nil {
		return nil, nil
	}

	channelDomain, err = convertor.ConvertEntitiesToDomains(channelEntity)
	if err != nil {
		return nil, err
	}
	return channelDomain, nil
}

func (repository *ChannelRepository) DeleteChannel(ctx context.Context, id string) error {
	err := repository.channelMongoService.DeleteChannel(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
