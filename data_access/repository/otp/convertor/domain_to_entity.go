package convertor

import (
	otpEntity "github.com/likhithkp/ping/data_access/mongo/otp"
	"github.com/likhithkp/ping/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDomainToEntity(domain domain.OtpDomain) (*otpEntity.OtpEntity, error) {

	var oid primitive.ObjectID
	var userId primitive.ObjectID
	var err error
	if domain.Id != "" {
		oid, err = primitive.ObjectIDFromHex(domain.Id)
		if err != nil {
			return nil, err
		}
	}

	if domain.UserId != "" {
		userId, err = primitive.ObjectIDFromHex(domain.UserId)
		if err != nil {
			return nil, err
		}
	}

	return &otpEntity.OtpEntity{
		Id:        oid,
		UserId:    userId,
		Email:     domain.Email,
		Otp:       domain.Otp,
		CreatedAt: domain.CreateAt,
		UpdatedAt: domain.UpdatedAt,
	}, nil
}
