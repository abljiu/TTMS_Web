package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type TheaterDao struct {
	*gorm.DB
}

func NewTheaterDao(ctx context.Context) *TheaterDao {
	return &TheaterDao{NewDBClient(ctx)}
}

func NewTheaterDaoByDB(db *gorm.DB) *TheaterDao {
	return &TheaterDao{db}
}

func (dao *TheaterDao) GetTheaterByTheaterID(id uint) (theater *model.Theater, err error) {
	err = dao.DB.Model(&model.Theater{}).Where("id=?", id).First(&theater).Error
	return
}
