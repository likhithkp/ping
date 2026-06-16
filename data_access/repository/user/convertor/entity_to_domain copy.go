package convertor

import (
	userEntity "github.com/likhithkp/ping/data_access/mongo/user"
	"github.com/likhithkp/ping/domain"
)

func ConvertEntityToDomain(entity *userEntity.UserEntity) *domain.UserDomain {
	return &domain.UserDomain{
		Id:          entity.Id.Hex(),
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		UserName:    entity.UserName,
		Bio:         entity.Bio,
		DateOfBirth: entity.DateOfBirth,
		Password:    entity.Password,
		Email:       entity.Email,
		PhoneNumber: entity.PhoneNumber,
	}
}
