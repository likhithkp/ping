package otp

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OtpMongoService struct {
	collection *mongo.Collection
}

func NewOtpMongoService(database *mongo.Database) *OtpMongoService {
	return &OtpMongoService{
		collection: database.Collection("otps"),
	}
}

func (service *OtpMongoService) UpsertOtp(ctx context.Context, otp *OtpEntity) error {
	if otp.Id == primitive.NilObjectID {
		otp.Id = primitive.NewObjectID()
	}

	filter := bson.M{"_id": otp.Id}
	update := bson.M{"$set": otp}
	opts := options.Update().SetUpsert(true)
	_, err := service.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (service *OtpMongoService) GetLatestOtpByEmail(ctx context.Context, email string) (otpEntity *OtpEntity, err error) {
	var otp OtpEntity
	filter := bson.M{"email": email}

	err = service.collection.FindOne(ctx, filter, options.FindOne().SetSort(bson.D{{Key: "createdAt", Value: -1}})).Decode(&otp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &otp, nil
}
