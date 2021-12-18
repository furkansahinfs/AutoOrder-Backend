package user_information

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetUserInformation(id int) (*model.UserInformation, error)
}

type Writer interface {
	StoreUserInformation(userInformation model.UserInformation) (int64, error)
	UpdateUserInformation(userInformation model.UserInformation, id int64) (int64, error)
	DeleteUserInformation(id int64) (int64, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
