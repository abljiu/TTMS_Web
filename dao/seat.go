package dao

import (
	"context"
	"gorm.io/gorm"
)

type SeatDao struct {
	*gorm.DB
}

func NewSeatDao(ctx context.Context) *SeatDao {
	return &SeatDao{NewDBClient(ctx)}
}

//func (dao *SeatDao) ModifySeat(product *model.ModifySeat) (err error) {
//	return dao.DB.Model(&model.Session{}).Where("id=?", product.HallID).Updates(model.ModifySeat{Seat: product.Seat}).Error
//}
