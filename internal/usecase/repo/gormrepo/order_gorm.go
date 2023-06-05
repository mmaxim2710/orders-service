package gormrepo

import (
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
	err := r.db.Model(&entity.Order{}).Create(order).First(newOrder).Error
	return newOrder, err
}
