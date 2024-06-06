package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Preload("Movie").Find(&carousel).Limit(6).Error
	return
}

func (dao *CarouselDao) CreateCarousel(item *model.Carousel) (err error) {
	err = dao.DB.Model(&model.Carousel{}).Create(&item).Error
	return
}
