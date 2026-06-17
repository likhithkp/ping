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

func (repository *UserRepository) UpdateUser(ctx context.Context, userDomain *domain.UserDomain) error {
	userEntity, err := convertor.ConvertDomainToEntity(*userDomain)
	if err != nil {
		return err
	}

	err = repository.userMongoService.UpdateUser(ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) InsertUser(ctx context.Context, userDomain *domain.UserDomain) error {
	userEntity, err := convertor.ConvertDomainToEntity(*userDomain)
	if err != nil {
		return err
	}

	err = repository.userMongoService.InsertUser(ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) GetUserById(ctx context.Context, id string) (userDomain *domain.UserDomain, err error) {
	userEntity, err := repository.userMongoService.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, nil
	}

	userDomain = convertor.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}

func (repository *UserRepository) GetUserByEmail(ctx context.Context, email string) (userDomain *domain.UserDomain, err error) {
	userEntity, err := repository.userMongoService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, nil
	}

	userDomain = convertor.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}

func (repository *UserRepository) GetUserByPhoneNumber(ctx context.Context, email string) (userDomain *domain.UserDomain, err error) {
	userEntity, err := repository.userMongoService.GetUserByPhoneNumber(ctx, email)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, nil
	}

	userDomain = convertor.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}
