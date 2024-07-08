package handler

type UserResponse struct {
	ProfilePicture string `json:"profile_picture"`
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
}
