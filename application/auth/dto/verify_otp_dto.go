package dto

type VerifyOtpRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
