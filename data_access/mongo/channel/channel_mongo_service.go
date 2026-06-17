package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChannelMongoService struct {
	collection *mongo.Collection
}

func NewChannelMongoService(database *mongo.Database) *ChannelMongoService {
	return &ChannelMongoService{
		collection: database.Collection("channels"),
	}
}

func (service *ChannelMongoService) InsertChannel(ctx context.Context, channel *ChannelEntity) error {
	_, err := service.collection.InsertOne(ctx, channel)
	if err != nil {
		return err
	}
	return nil
}

func (service *ChannelMongoService) UpdateChannel(ctx context.Context, channel *ChannelEntity) error {
	filter := bson.M{"_id": channel.Id}
	update := bson.M{"$set": channel}
	_, err := service.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (service *ChannelMongoService) GetChannelById(ctx context.Context, id string) (*ChannelEntity, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var channel ChannelEntity
	filter := bson.M{"_id": oid}

	err = service.collection.FindOne(ctx, filter).Decode(&channel)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &channel, nil
}

func (service *ChannelMongoService) GetChannelsByUserId(ctx context.Context, id string) ([]*ChannelEntity, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"users.userId": oid,
	}

	cursor, err := service.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var channels []*ChannelEntity
	for cursor.Next(ctx) {
		var channel ChannelEntity
		err := cursor.Decode(&channel)
		if err != nil {
			return nil, err
		}
		channels = append(channels, &channel)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}

func (service *ChannelMongoService) DeleteChannel(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}
	_, err = service.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
