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

func (userRepository *UserRepository) GetUserByEmail(ctx context.Context, email string) (userDomain *domain.UserDomain, err error) {
	userEntity, err := userRepository.userMongoService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, nil
	}

	userDomain = convertor.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}

func (userRepository *UserRepository) GetUserByPhoneNumber(ctx context.Context, email string) (userDomain *domain.UserDomain, err error) {
	userEntity, err := userRepository.userMongoService.GetUserByPhoneNumber(ctx, email)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, nil
	}

	userDomain = convertor.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}
