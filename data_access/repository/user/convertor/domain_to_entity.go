package convertor

import (
	userEntity "github.com/likhithkp/ping/data_access/mongo/user"
	"github.com/likhithkp/ping/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDomainToEntity(domain domain.UserDomain) (*userEntity.UserEntity, error) {

	var oid primitive.ObjectID
	var err error
	if domain.Id != "" {
		oid, err = primitive.ObjectIDFromHex(domain.Id)
		if err != nil {
			return nil, err
		}
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
