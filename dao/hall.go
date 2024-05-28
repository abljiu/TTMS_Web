package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type HallDao struct {
	*gorm.DB
}

func NewHallDao(ctx context.Context) *HallDao {
	return &HallDao{NewDBClient(ctx)}
}

func NewHallDaoByDB(db *gorm.DB) *HallDao {
	return &HallDao{db}
}

func (dao *HallDao) GetHallByHallID(id uint) (hall *model.Hall, err error) {
	err = dao.DB.Model(&model.Hall{}).Where("id=?", id).First(&hall).Error
	return
}
