package otp

import (
	"context"

	otpMongoService "github.com/likhithkp/ping/data_access/mongo/otp"
	"github.com/likhithkp/ping/data_access/repository/otp/convertor"
	"github.com/likhithkp/ping/domain"
)

type OtpRepository struct {
	otpMongoService *otpMongoService.OtpMongoService
}

func NewOtpRepository(otpMongoService *otpMongoService.OtpMongoService) *OtpRepository {
	return &OtpRepository{
		otpMongoService: otpMongoService,
	}
}

func (otpRepository *OtpRepository) UpsertOtp(ctx context.Context, otpDomain *domain.OtpDomain) error {
	otpEntity, err := convertor.ConvertDomainToEntity(*otpDomain)
	if err != nil {
		return err
	}

	err = otpRepository.otpMongoService.UpsertOtp(ctx, otpEntity)
	if err != nil {
		return err
	}
	return nil
}

func (otpRepository *OtpRepository) GetLatestOtpByEmail(ctx context.Context, email string) (otpDomain *domain.OtpDomain, err error) {
	otpEntity, err := otpRepository.otpMongoService.GetLatestOtpByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if otpEntity == nil {
		return nil, nil
	}

	otpDomain = convertor.ConvertEntityToDomain(otpEntity)
	return otpDomain, nil
}
