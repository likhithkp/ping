package dto

type SignUpRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	UserName    string `json:"userName"`
	Bio         string `json:"bio"`
	DateOfBirth string `json:"dateOfBirth"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
