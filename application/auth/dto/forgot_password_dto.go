package dto

import _const "github.com/likhithkp/ping/utils/const"

type ForgotPasswordRequest struct {
	PhoneNumber    string                `json:"phoneNumber"`
	Email          string                `json:"email"`
	IdentifierType _const.IdentifierType `json:"identifierType"`
}
