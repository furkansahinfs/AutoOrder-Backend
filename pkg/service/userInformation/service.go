package userinformation

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}
