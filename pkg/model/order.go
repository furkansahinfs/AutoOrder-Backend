package model

type Order struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Date     string `json:"date"`
	ItemName string `json:"item_name"`
	Count    int    `json:"count"`
}
