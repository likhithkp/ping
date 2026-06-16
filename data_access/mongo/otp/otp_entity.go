package otp

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OtpEntity struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"userId"`
	Email     string             `bson:"email"`
	Otp       string             `bson:"otp"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
