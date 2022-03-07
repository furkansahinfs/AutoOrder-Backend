package orders

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetOrders(userID int64) ([]*model.Order, error)
	GetOrder(id int64, userID int64) (*model.Order, error)
}

type Writer interface {
	CreateOrder(order *model.Order, userID int64) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
