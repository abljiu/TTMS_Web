package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) AddOrder(order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&order).Error
}

func (dao *OrderDao) GetOrderByID(id uint) (order *model.Order, err error) {
	err = dao.DB.Preload("Movie").Preload("Theater").Preload("Hall").Preload("Session").Model(&model.Order{}).Where("id=?", id).First(&order).Error
	return
}

// todo 根据座位找订单id

func (dao *OrderDao) GetOrderIDBySeat(seat string) (id uint, err error) {
	err = dao.DB.Model(&model.Order{}).Where("seat like ?").Error
	return
}

func (dao *OrderDao) CheckOrderTypeByID(id uint) (status uint, err error) {
	order := &model.Order{}
	err = dao.DB.Model(&model.Order{}).Where("id=?", id).First(&order).Error
	return order.Type, err
}
