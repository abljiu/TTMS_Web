package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
	"sync"
	"time"
)

type HallDao struct {
	*gorm.DB
	mu sync.RWMutex
}

func NewHallDao(ctx context.Context) *HallDao {
	return &HallDao{DB: NewDBClient(ctx)}
}

func NewHallDaoByDB(db *gorm.DB) *HallDao {
	return &HallDao{DB: db}
}

func (dao *HallDao) CountHallByTheaterID(theater uint) (total int64, err error) {
	dao.mu.RLock()
	defer dao.mu.RUnlock()
	err = dao.DB.Model(&model.Hall{}).Where("theater_id=?", theater).Count(&total).Error
	return
}

func (dao *HallDao) ListHallByTheaterID(theater uint, page model.BasePage) (products []*model.Hall, err error) {
	dao.mu.RLock()
	defer dao.mu.RUnlock()
	err = dao.DB.Model(&model.Hall{}).Where("theater_id=? and deleted_at is NULL", theater).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *HallDao) CreateHall(product *model.Hall) (err error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	return dao.DB.Model(&model.Hall{}).Create(&product).Error
}

func (dao *HallDao) DeleteHall(product *model.Hall) (err error) {
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.ID).Delete(&product).Error
}

func (dao *HallDao) UpdateHall(product *model.Hall) (err error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	product.CreatedAt = time.Now()
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.ID).Save(&product).Error
}

func (dao *HallDao) GetHall(product *model.Hall) (err error) {
	return dao.DB.Model(&model.Hall{}).Where("id=?", product.ID).First(&product).Error
}

func (dao *HallDao) GetHallByHallID(id uint) (hall *model.Hall, err error) {
	err = dao.DB.Model(&model.Hall{}).Where("id=?", id).First(&hall).Error
	return
}
