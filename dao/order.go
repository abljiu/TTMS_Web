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

func (dao *OrderDao) AddOrder(order *model.Order) (*model.Order, error) {
	if err := dao.DB.Model(&model.Order{}).Create(&order).Error; err != nil {
		return nil, err
	} else {
		return order, nil
	}
}

func (dao *OrderDao) GetOrderByOrderID(id uint) (order *model.Order, err error) {
	err = dao.DB.Preload("Movie").Preload("Theater").Preload("Session").Model(&model.Order{}).Where("id=?", id).First(&order).Error
	return
}

func (dao *OrderDao) UpdateOrderByID(id uint, order *model.Order) error {
	err := dao.DB.Model(&model.Order{}).Where("id=?", id).Updates(&order).Error
	return err
}

//func (dao *OrderDao) GetOrderIDBySeat(seat string) (id uint, err error) {
//	err = dao.DB.Model(&model.Order{}).Where("seat like ?"sea).Error
//	return
//}

func (dao *OrderDao) CheckOrderTypeByID(id uint) (status uint, err error) {
	order := &model.Order{}
	err = dao.DB.Model(&model.Order{}).Where("id=?", id).First(&order).Error
	return order.Type, err
}

func (dao *OrderDao) IsUserBuyMovie(userId, movieId uint) (bool, error) {
	var count int64
	err := dao.DB.Model(&model.Order{}).Where("user_id = ? AND movie_id = ?", userId, movieId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (dao *OrderDao) DeleteOrderByID(uid uint) error {
	err := dao.DB.Where("id=?", uid).Delete(&model.Order{}).Error
	return err
}

func (dao *OrderDao) ListUserOrders(userId uint) (orders []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Preload("Movie").Preload("Theater").Preload("Session").Where("user_id = ?", userId).Find(&orders).Error

	return
}

func (dao *OrderDao) ListOrdersBySessionID(sessionId uint) (orders []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("session_id = ?", sessionId).Find(&orders).Error

	return
}
