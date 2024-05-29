package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
	"time"
)

type HallDao struct {
	*gorm.DB
}

func NewHallDao(ctx context.Context) *HallDao {
	return &HallDao{NewDBClient(ctx)}
}

func (dao *HallDao) CountHallByTheaterID(theater int) (total int64, err error) {
	err = dao.DB.Model(&model.Hall{}).Where("theater_id=?", theater).Count(&total).Error
	return
}
func NewHallDaoByDB(db *gorm.DB) *HallDao {
	return &HallDao{db}
}

func (dao *HallDao) ListHallByTheaterID(theater int, page model.BasePage) (products []*model.Hall, err error) {
	err = dao.DB.Where("theater_id=? and deleted_at is NULL", theater).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *HallDao) CreateHall(product *model.Hall) (err error) {
	return dao.DB.Model(&model.Hall{}).Create(&product).Error
}

func (dao *HallDao) DeleteHall(product *model.Hall) (err error) {
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.ID).Delete(&product).Error
}

func (dao *HallDao) UpdateHall(product *model.Hall) (err error) {
	product.CreatedAt = time.Now()
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.ID).Save(&product).Error
}

func (dao *HallDao) GetHall(product *model.Hall) (err error) {
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.ID).First(&product).Error
}
