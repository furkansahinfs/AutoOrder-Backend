package orders

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
	"github.com/google/uuid"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) SaveOrders(order []model.OrderResponse, userID int64, path string) error {
	id := uuid.New().String()

	for i, _ := range order {
		order[i].ImagePath = path
	}

	return s.repository.GetOrdersRepository().SaveOrder(order, userID, id)
}

func (s *Service) GetOrders(userID int64) ([]model.OrderHistory, error) {
	history, err := s.repository.GetOrdersRepository().GetOrdersWithGroupByOrderID(userID)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (s *Service) GetOrder(orderID string, userID int64) (model.OrderHistory, error) {
	history, err := s.repository.GetOrdersRepository().GetOrder(userID, orderID)
	if err != nil {
		return model.OrderHistory{}, err
	}
	return history, nil
}
