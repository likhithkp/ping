package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	Id          primitive.ObjectID `bson:"_id"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	UserName    string             `bson:"userName"`
	Bio         string             `bson:"bio"`
	DateOfBirth time.Time          `bson:"dateOfBirth"`
	Image       string             `bson:"image"`
	Password    string             `bson:"password"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phoneNumber"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}
