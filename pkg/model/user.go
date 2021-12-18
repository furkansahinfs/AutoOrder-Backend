package model

//User data
type User struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	FullName          string `json:"fullName"`
	Token             string `json:"token"`
	UserInformationID int64  `json:"user_information_id"`
}
