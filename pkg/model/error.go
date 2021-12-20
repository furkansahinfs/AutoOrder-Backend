package model

//Error data
type ErrorJson []struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
