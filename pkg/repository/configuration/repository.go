package configuration

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetConfiguration(id int64, item_type string) ([]model.Item, error)
}

type Writer interface {
	DeleteConfiguration(items []model.Item, item_type string, user_id int64) error
	UpdateConfiguration(items []model.Item, item_type string, user_id int64) error
	StoreConfiguration(items []model.Item, user_id int64) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
