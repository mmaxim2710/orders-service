package usecase

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"gorm.io/gorm"
)

const (
	CREATED string = "CREATED"
)

type OrderUseCase struct {
	serviceRepo ServiceRepo
	orderRepo   OrderRepo
}

func NewOrderUseCase(serviceRepo ServiceRepo, orderRepo OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		serviceRepo: serviceRepo,
		orderRepo:   orderRepo,
	}
}

func (u *OrderUseCase) Create(serviceIDs []uuid.UUID) ([]*entity.Order, error) {
	if len(serviceIDs) < 1 {
		return nil, ErrEmptySlice
	}

	orders := make([]*entity.Order, 0, len(serviceIDs))
	orderID := uuid.New()

	for _, sID := range serviceIDs {
		isExists, err := u.serviceRepo.IsServiceExists(sID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if !isExists {
			return nil, ErrServiceNotExists
		}

		order := &entity.Order{
			OrderID:   orderID,
			ServiceID: sID,
			Status:    CREATED,
		}

		newOrder, err := u.orderRepo.Create(order)
		if err != nil {
			return nil, err
		}

		orders = append(orders, newOrder)
	}

	return orders, nil
}
