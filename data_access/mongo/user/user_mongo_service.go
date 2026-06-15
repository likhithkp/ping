package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": user,
	}
	opts := options.Update().SetUpsert(true)
	_, err := userMongoService.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
