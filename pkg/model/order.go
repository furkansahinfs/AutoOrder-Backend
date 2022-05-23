package model

type OrderResponse struct {
	Name      string `json:"name"`
	Brand     int64  `json:"brand"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	ID        int64  `json:"id"`
	ImagePath string `json:"image_path"`
}

type OrderResponseJava struct {
	OrderResponse []OrderResponse `json:"data"`
}
type OrderRequestJava struct {
	OrderRequest string `json:"orderItems"`
}

type OrderHistory struct {
	ID        int64   `json:"id"`
	Orders    []Order `json:"orders"`
	ImagePath string  `json:"image_path"`
}

type Order struct {
	Name     string `json:"name"`
	Brand    int64  `json:"brand"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
