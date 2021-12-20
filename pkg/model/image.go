package model

//UserInformation data
type ImageData struct {
	Id     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Path   string `json:"path"`
}
