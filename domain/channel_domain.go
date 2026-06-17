package domain

import (
	"time"
)

type User struct {
	UserId    string
	Username  string
	Bio       string
	Image     string
	IsBlocked bool
}

type ChannelDomain struct {
	Id             string
	Users          []User
	LastMessage    string
	LastMessagedAt time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
