package dto

type GetDetailsRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	UserName    string `json:"userName"`
	Bio         string `json:"bio"`
	Image       string `json:"image"`
	DateOfBirth string `json:"dateOfBirth"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
