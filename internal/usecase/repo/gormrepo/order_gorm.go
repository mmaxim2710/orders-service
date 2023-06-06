package gormrepo

import (
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/entity"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
	l  logger.Interface
}

func NewOrderRepo(db *gorm.DB, l logger.Interface) *OrderRepo {
	return &OrderRepo{
		db: db,
		l:  l,
	}
}

func (r *OrderRepo) Create(order *entity.Order) (*entity.Order, error) {
	newOrder := &entity.Order{}
	err := r.db.Model(&entity.Order{}).
		Create(order).
		First(newOrder).Error
	return newOrder, err
}

func (r *OrderRepo) GetByOrderID(orderID uuid.UUID) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Model(&entity.Order{}).
		Where("order_id = ?", orderID).
		Find(&orders).Error
	return orders, err
}

func (r *OrderRepo) IsOrderExists(orderID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Order{}).Where("order_id = ?", orderID).Count(&count).Error
	return count > 0, err
}
