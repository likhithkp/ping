package dto

import _const "github.com/likhithkp/ping/utils/const"

type SignInRequest struct {
	PhoneNumber    string                `json:"phoneNumber"`
	Email          string                `json:"email"`
	Password       string                `json:"password"`
	IdentifierType _const.IdentifierType `json:"identifierType"`
}
