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

func (dao *TheaterDao) CreateTheater(Theater *model.Theater) error {
	return dao.DB.Model(&model.Theater{}).Create(&Theater).Error
}

func (dao *TheaterDao) GetTheaterByID(id uint) (Theater *model.Theater, err error) {
	err = dao.DB.Model(&model.Theater{}).Where("id=?", id).First(&Theater).Error
	return
}

func (dao *TheaterDao) UpdateTheaterByID(uid uint, Theater *model.Theater) error {
	err := dao.DB.Model(&model.Theater{}).Where("id=?", uid).Updates(&Theater).Error
	return err
}

func (dao *TheaterDao) DeleteTheaterByID(id uint) (Theater *model.Theater, err error) {
	err = dao.DB.Model(&model.Theater{}).Where("id=?", id).Delete(&Theater).Error
	return
}
func (dao *TheaterDao) CountTheater() (total int64, err error) {
	err = dao.DB.Model(&model.Theater{}).Count(&total).Error
	return
}

func (dao *TheaterDao) ListTheater(page model.BasePage) (Theaters []*model.Theater, err error) {

	err = dao.DB.Model(&model.Theater{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&Theaters).Error
	return
}
func (dao *TheaterDao) SearchTheater(info string, page model.BasePage) (products []*model.Theater, err error) {
	err = dao.DB.Model(&model.Theater{}).
		Where("name LIKE ?", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}
