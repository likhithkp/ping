package dto

type EditProfileRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	UserName    string `json:"userName"`
	Bio         string `json:"bio"`
	NewImage    string `json:"newImage"`
	DateOfBirth string `json:"dateOfBirth"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
