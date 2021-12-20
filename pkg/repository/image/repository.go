package image

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/model"

type Reader interface {
	GetImagePath(id int64) (*model.ImageData, error)
}

type Writer interface {
	StoreImagePath(baseImagePath string, user_id int64) (int64, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
