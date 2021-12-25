package configuration

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
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

func (s *Service) GetUserConfiguration(id int64, item_type string) ([]model.Item, error) {
	configurations, err := s.repository.GetConfigurationRepository().GetConfiguration(id, item_type)
	if err != nil {
		return nil, nil
	}
	return configurations, nil
}
