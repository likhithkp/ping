package user

import (
	"context"

	userMongoService "github.com/likhithkp/ping/data_access/mongo/user"
	"github.com/likhithkp/ping/data_access/repository/user/convertor"
	"github.com/likhithkp/ping/domain"
)

type UserRepository struct {
	userMongoService *userMongoService.UserMongoService
}

func NewUserRepository(userMongoService *userMongoService.UserMongoService) *UserRepository {
	return &UserRepository{
		userMongoService: userMongoService,
	}
}

func (userRepository *UserRepository) UpsertUser(ctx context.Context, userDomain *domain.UserDomain) error {
	userEntity, err := convertor.ConvertDomainToEntity(*userDomain)
	if err != nil {
		return err
	}

	err = userRepository.userMongoService.UpsertUser(ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}
