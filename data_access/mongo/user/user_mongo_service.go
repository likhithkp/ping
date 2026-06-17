package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoService struct {
	collection *mongo.Collection
}

func NewUserMongoService(database *mongo.Database) *UserMongoService {
	return &UserMongoService{
		collection: database.Collection("users"),
	}
}

func (service *UserMongoService) InsertUser(ctx context.Context, user *UserEntity) error {
	_, err := service.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserMongoService) UpdateUser(ctx context.Context, user *UserEntity) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}
	_, err := service.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserMongoService) GetUserById(ctx context.Context, id string) (userEntity *UserEntity, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user UserEntity
	filter := bson.M{"_id": oid}

	err = service.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (service *UserMongoService) GetUserByEmail(ctx context.Context, email string) (userEntity *UserEntity, err error) {
	var user UserEntity
	filter := bson.M{"email": email}

	err = service.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (service *UserMongoService) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userEntity *UserEntity, err error) {
	var user UserEntity
	filter := bson.M{"phoneNumber": phoneNumber}

	err = service.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (service *UserMongoService) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*UserEntity, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range userIDs {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, oid)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	cursor, err := service.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*UserEntity
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
