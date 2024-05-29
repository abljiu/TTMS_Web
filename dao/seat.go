package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type SeatDao struct {
	*gorm.DB
}

func NewSeatDao(ctx context.Context) *SeatDao {
	return &SeatDao{NewDBClient(ctx)}
}

func (dao *SeatDao) ModifySeat(product *model.ModifySeat) (err error) {
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.HallID).Updates(model.ModifySeat{Seat: product.Seat}).Error
}
