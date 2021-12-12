package user

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetUser(user model.User) (*model.User, error)
	CheckExist(user model.User) (bool, error)
}

type Writer interface {
	StoreUser(user model.User) (*model.User, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
