package model

type Order struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	OrderDate string `json:"order_date"`
	Products  []Item
}
