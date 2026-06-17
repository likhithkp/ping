package convertor

import (
	"time"

	"github.com/likhithkp/ping/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertCreateChannelDtoToDomain(usersDomains []*domain.UserDomain) *domain.ChannelDomain {
	var users []domain.User
	for _, user := range usersDomains {
		users = append(users, domain.User{
			UserId:    user.Id,
			Username:  user.UserName,
			Bio:       user.Bio,
			Image:     user.Image,
			IsBlocked: false,
		})
	}

	return &domain.ChannelDomain{
		Id:        primitive.NewObjectID().Hex(),
		Users:     users,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
