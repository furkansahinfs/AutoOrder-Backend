package orders

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetOrdersWithGroupByOrderID(userID int64) ([]model.OrderHistory, error)
	GetOrder(userID int64, orderID string) (model.OrderHistory, error)
}

type Writer interface {
	SaveOrder(order []model.OrderResponse, userID int64, orderID string) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
