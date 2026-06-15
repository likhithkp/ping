package convertor

import (
	"fmt"

	userEntity "github.com/likhithkp/ping/data_access/mongo/user"
	"github.com/likhithkp/ping/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDomainToEntity(domain domain.UserDomain) (*userEntity.UserEntity, error) {

	oid, err := primitive.ObjectIDFromHex(domain.Id)
	if err != nil {
		fmt.Printf("[domain_to_entity.go] %s ", err.Error())
		return nil, err
	}

	return &userEntity.UserEntity{
		Id:          oid,
		FirstName:   domain.FirstName,
		LastName:    domain.LastName,
		UserName:    domain.UserName,
		Bio:         domain.Bio,
		DateOfBirth: domain.DateOfBirth,
		Password:    domain.Password,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		CreatedAt:   domain.CreateAt,
		UpdatedAt:   domain.UpdatedAt,
	}, nil
}
