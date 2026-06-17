package convertor

import (
	"github.com/likhithkp/ping/application/channel/dto"
	"github.com/likhithkp/ping/domain"
)

func ConvertChannelListToDto(channelDomain []*domain.ChannelDomain) []*dto.ChannelListDto {
	var channelList []*dto.ChannelListDto

	for _, c := range channelDomain {
		var users []dto.User
		for _, u := range c.Users {
			user := dto.User{
				UserId:    u.UserId,
				Username:  u.Username,
				Bio:       u.Bio,
				Image:     u.Image,
				IsBlocked: u.IsBlocked,
			}
			users = append(users, user)
		}

		channel := &dto.ChannelListDto{
			Id:             c.Id,
			Users:          users,
			LastMessage:    c.LastMessage,
			LastMessagedAt: c.LastMessagedAt,
		}

		channelList = append(channelList, channel)
	}

	return channelList
}

func ConvertChannelDetailsToDto(channelDomain *domain.ChannelDomain) *dto.ChannelListDto {
	var users []dto.User

	for _, u := range channelDomain.Users {
		user := dto.User{
			UserId:    u.UserId,
			Username:  u.Username,
			Bio:       u.Bio,
			Image:     u.Image,
			IsBlocked: u.IsBlocked,
		}
		users = append(users, user)
	}

	channel := &dto.ChannelListDto{
		Id:             channelDomain.Id,
		Users:          users,
		LastMessage:    channelDomain.LastMessage,
		LastMessagedAt: channelDomain.LastMessagedAt,
	}

	return channel
}
