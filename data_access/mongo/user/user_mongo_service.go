package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoService struct {
	collection *mongo.Collection
}

func NewUserMongoService(database *mongo.Database) *UserMongoService {
	return &UserMongoService{
		collection: database.Collection("users"),
	}
}

func (userMongoService *UserMongoService) UpsertUser(ctx context.Context, user *UserEntity) error {
	if user.Id == primitive.NilObjectID {
		user.Id = primitive.NewObjectID()
	}

	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}
	opts := options.Update().SetUpsert(true)
	_, err := userMongoService.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (userMongoService *UserMongoService) GetUserById(ctx context.Context, id string) (userEntity *UserEntity, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user UserEntity
	filter := bson.M{"_id": oid}

	err = userMongoService.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (userMongoService *UserMongoService) GetUserByEmail(ctx context.Context, email string) (userEntity *UserEntity, err error) {
	var user UserEntity
	filter := bson.M{"email": email}

	err = userMongoService.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (userMongoService *UserMongoService) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userEntity *UserEntity, err error) {
	var user UserEntity
	filter := bson.M{"phoneNumber": phoneNumber}

	err = userMongoService.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
