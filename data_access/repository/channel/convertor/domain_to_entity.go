package convertor

import (
	channelEntity "github.com/likhithkp/ping/data_access/mongo/channel"
	"github.com/likhithkp/ping/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDomainToEntity(domain domain.ChannelDomain) (*channelEntity.ChannelEntity, error) {
	var oid primitive.ObjectID
	var err error

	if domain.Id != "" {
		oid, err = primitive.ObjectIDFromHex(domain.Id)
		if err != nil {
			return nil, err
		}
	}

	var users []channelEntity.User
	for _, u := range domain.Users {
		userOid, err := primitive.ObjectIDFromHex(u.UserId)
		if err != nil {
			return nil, err
		}

		user := channelEntity.User{
			UserId:    userOid,
			Username:  u.Username,
			Bio:       u.Bio,
			Image:     u.Image,
			IsBlocked: u.IsBlocked,
		}
		users = append(users, user)
	}

	return &channelEntity.ChannelEntity{
		Id:             oid,
		Users:          users,
		LastMessage:    domain.LastMessage,
		LastMessagedAt: domain.LastMessagedAt,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}, nil
}
