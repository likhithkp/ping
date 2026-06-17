package convertor

import (
	"time"

	"github.com/likhithkp/ping/application/user/dto"
	"github.com/likhithkp/ping/domain"
)

func ConvertDomainToGetDetails(userDomain *domain.UserDomain) *dto.GetDetailsRequest {
	return &dto.GetDetailsRequest{
		FirstName:   userDomain.FirstName,
		LastName:    userDomain.LastName,
		UserName:    userDomain.UserName,
		Bio:         userDomain.Bio,
		Image:       userDomain.Image,
		DateOfBirth: userDomain.DateOfBirth.Format(time.RFC3339),
		Email:       userDomain.Email,
		PhoneNumber: userDomain.PhoneNumber,
	}
}
