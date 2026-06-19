package domain

import (
	"time"
)

type UserDomain struct {
	Id          string
	FirstName   string
	LastName    string
	UserName    string
	Bio         string
	Image       string
	DateOfBirth time.Time
	Password    string
	Email       string
	PhoneNumber string
	LastSeen    time.Time
	CreateAt    time.Time
	UpdatedAt   time.Time
}
