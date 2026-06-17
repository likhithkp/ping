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
		Image:       entity.Image,
		DateOfBirth: entity.DateOfBirth,
		Password:    entity.Password,
		Email:       entity.Email,
		PhoneNumber: entity.PhoneNumber,
	}
}

func ConvertEntitiesToDomain(entities []*userEntity.UserEntity) []*domain.UserDomain {
	var domains []*domain.UserDomain
	for _, entity := range entities {
		domains = append(domains, &domain.UserDomain{
			Id:          entity.Id.Hex(),
			FirstName:   entity.FirstName,
			LastName:    entity.LastName,
			UserName:    entity.UserName,
			Bio:         entity.Bio,
			Image:       entity.Image,
			DateOfBirth: entity.DateOfBirth,
			Password:    entity.Password,
			Email:       entity.Email,
			PhoneNumber: entity.PhoneNumber,
		})
	}
	return domains
}
