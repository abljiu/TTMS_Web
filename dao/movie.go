package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type MovieDao struct {
	*gorm.DB
}

func NewMovieDao(ctx context.Context) *MovieDao {
	return &MovieDao{NewDBClient(ctx)}
}

func NewMovieDaoByDB(db *gorm.DB) *MovieDao {
	return &MovieDao{db}
}

func (dao *MovieDao) CreateMovie(product *model.Movie) (err error) {
	return dao.DB.Model(&model.Movie{}).Create(&product).Error
}

func (dao *MovieDao) CountMovieByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Movie{}).Where(condition).Count(&total).Error
	return
}

func (dao *MovieDao) ListMovieByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Movie, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *MovieDao) SearchMovie(info string, page model.BasePage) (products []*model.Movie, err error) {
	err = dao.DB.Model(&model.Movie{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}
