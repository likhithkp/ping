package convertor

import (
	"time"

	"github.com/likhithkp/ping/application/user/dto"
	"github.com/likhithkp/ping/domain"
)

func ConvertEditProfileDtoToDomain(id string, requestBody *dto.EditProfileRequest) (*domain.UserDomain, error) {
	dob, err := time.Parse("2006-01-02", requestBody.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &domain.UserDomain{
		Id:          id,
		FirstName:   requestBody.FirstName,
		LastName:    requestBody.LastName,
		UserName:    requestBody.UserName,
		Bio:         requestBody.Bio,
		DateOfBirth: dob,
		Email:       requestBody.Email,
		PhoneNumber: requestBody.PhoneNumber,
		UpdatedAt:   time.Now().UTC(),
	}, nil
}
