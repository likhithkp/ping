package convertor

import (
	"time"

	"github.com/likhithkp/ping/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertOtpReqToDomain(email, userId, otp string) *domain.OtpDomain {
	return &domain.OtpDomain{
		Id:        primitive.NewObjectID().Hex(),
		UserId:    userId,
		Email:     email,
		Otp:       otp,
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
