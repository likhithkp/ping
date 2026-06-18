package convertor

import (
	channelEntity "github.com/likhithkp/ping/data_access/mongo/channel"
	"github.com/likhithkp/ping/domain"
)

func ConvertEntityToDomain(entity *channelEntity.ChannelEntity) (*domain.ChannelDomain, error) {
	channelID := entity.Id.Hex()

	var users []domain.User
	for _, u := range entity.Users {
		user := domain.User{
			UserId:    u.UserId.Hex(),
			Username:  u.Username,
			Bio:       u.Bio,
			Image:     u.Image,
			IsBlocked: u.IsBlocked,
		}
		users = append(users, user)
	}

	return &domain.ChannelDomain{
		Id:             channelID,
		Users:          users,
		LastMessage:    entity.LastMessage,
		LastMessagedAt: entity.LastMessagedAt,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}, nil
}

func ConvertEntitiesToDomains(entities []*channelEntity.ChannelEntity) ([]*domain.ChannelDomain, error) {
	var domains []*domain.ChannelDomain

	for _, entity := range entities {
		domain, err := ConvertEntityToDomain(entity)
		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	return domains, nil
}
