package configuration

import (
	"errors"
	"strings"

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
		return nil, err
	}
	if len(configurations) > 0 {
		return configurations, nil
	}
	return nil, errors.New("User dont have a configuration")
}

func (s *Service) StoreUserConfiguration(id int64, item_type string, uitems []model.Item, allItems []model.Item) error {
	items, err := s.repository.GetConfigurationRepository().GetConfiguration(id, item_type)
	if err != nil {
		return err
	}
	if len(items) > 0 {
		return errors.New("User already have an " + item_type + " configuration ")
	}
	for index, item := range uitems {
		for _, enumItem := range allItems {
			if strings.EqualFold(item.Name, enumItem.Name) {
				uitems[index].Size = enumItem.Size
				uitems[index].Type = enumItem.Type
			}
		}
	}
	err = s.repository.GetConfigurationRepository().StoreConfiguration(uitems, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUserConfiguration(id int64, item_type string, items []model.Item, allItems []model.Item) error {
	items, err := s.repository.GetConfigurationRepository().GetConfiguration(id, item_type)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return errors.New("User dont have have an " + item_type + " configuration ")
	}
	for index, item := range items {
		for _, enumItem := range allItems {
			if item.Name == enumItem.Name {
				items[index].Size = enumItem.Size
				items[index].Type = enumItem.Type
			}
		}
	}
	err = s.repository.GetConfigurationRepository().UpdateConfiguration(id, items, item_type)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserConfiguration(id int64, item_type string) error {
	items, err := s.repository.GetConfigurationRepository().GetConfiguration(id, item_type)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return errors.New("User dont have have an " + item_type + " configuration ")
	}

	err = s.repository.GetConfigurationRepository().DeleteConfiguration(id, item_type)
	if err != nil {
		return err
	}
	return nil
}
