package domain

import (
	"time"
)

type OtpDomain struct {
	Id        string
	UserId    string
	Email     string
	Otp       string
	CreateAt  time.Time
	UpdatedAt time.Time
}
