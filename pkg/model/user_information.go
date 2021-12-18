package model

//UserInformation data
type UserInformation struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Id       int64  `json:"id"`
}
