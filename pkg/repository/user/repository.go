package user

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetUser(user model.User) (*model.User, error)
}

type Writer interface {
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
