package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId    primitive.ObjectID `bson:"userId"`
	Username  string             `bson:"username"`
	Bio       string             `bson:"bio"`
	Image     string             `bson:"image"`
	IsBlocked bool               `bson:"isBlocked"`
}

type ChannelEntity struct {
	Id             primitive.ObjectID `bson:"_id"`
	Users          []User             `bson:"users"`
	LastMessage    string             `bson:"lastMessage"`
	LastMessagedAt time.Time          `bson:"lastMessagedAt"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
