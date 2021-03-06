package model

type OrderResponse struct {
	Name      string  `json:"productName"`
	Brand     string  `json:"brandName"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	ID        int64   `json:"id"`
	ImagePath string  `json:"image_path"`
}

type OrderResponseJava struct {
	OrderResponse []OrderResponse `json:"data"`
}
type OrderRequestJava struct {
	OrderRequest string `json:"orderItems"`
}

type OrderHistory struct {
	ID        string  `json:"id"`
	Orders    []Order `json:"orders"`
	ImagePath string  `json:"image_path"`
}

type Order struct {
	Name      string  `json:"productName"`
	Brand     string  `json:"brandName"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	ImagePath string  `json:"image_path"`
}
