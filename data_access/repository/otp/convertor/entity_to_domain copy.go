package convertor

import (
	otpEntity "github.com/likhithkp/ping/data_access/mongo/otp"
	"github.com/likhithkp/ping/domain"
)

func ConvertEntityToDomain(entity *otpEntity.OtpEntity) *domain.OtpDomain {
	return &domain.OtpDomain{
		Id:        entity.Id.Hex(),
		UserId:    entity.UserId.Hex(),
		Email:     entity.Email,
		Otp:       entity.Otp,
		CreateAt:  entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
