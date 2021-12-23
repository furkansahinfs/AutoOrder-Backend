package image

import (
	"errors"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) SaveImagePath(path string, user_id int64) (int64, error) {
	id, err := s.repository.GetImageRepository().StoreImagePath(path, user_id)
	if err != nil {
		return -1, err
	}
	if id == -1 {
		return -1, errors.New("Error when saveImagePath")
	}
	return id, nil
}
