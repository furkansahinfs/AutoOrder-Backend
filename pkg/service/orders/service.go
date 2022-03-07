package orders

import (
	"strings"
	"time"

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

func (s *Service) GetOrders(userID int64) ([]*model.Order, error) {
	return s.repository.GetOrdersRepository().GetOrders(userID)
}

func (s *Service) GetOrder(id, userID int64) (*model.Order, error) {
	return s.repository.GetOrdersRepository().GetOrder(id, userID)
}

func (s *Service) CreateOrders(itemsNameString string, userID int64) error {
	// split itemNameString to array
	itemsName := strings.Split(itemsNameString, ",")
	for _, itemName := range itemsName {
		order := &model.Order{
			ItemName: itemName,
			UserID:   userID,
			Count:    1,
			Date:     time.Now().Format("2006-01-02"),
		}
		err := s.repository.GetOrdersRepository().CreateOrder(order, userID)
		if err != nil {
			return err
		}
	}
	return nil
}
