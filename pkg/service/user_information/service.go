package user_information

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

func (s *Service) StoreUserInformation(userInformation model.UserInformation) (int64, error) {
	id, err := s.repository.GetUserInformationRepository().StoreUserInformation(userInformation)
	if err != nil {
		return -1, nil
	}
	return id, nil
}

func (s *Service) DeleteUserInformation(id int64) (int64, error) {
	id, err := s.repository.GetUserInformationRepository().DeleteUserInformation(id)
	if err != nil {
		return -1, nil
	}
	return id, nil
}

func (s *Service) UpdateUserInformation(userInformation model.UserInformation, id int64) (int64, error) {
	id, err := s.repository.GetUserInformationRepository().UpdateUserInformation(userInformation, id)
	if err != nil {
		return -1, nil
	}
	return id, nil
}
