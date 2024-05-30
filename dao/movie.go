package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
	"strconv"
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

func (dao *MovieDao) CountMovieByCondition(categoryId uint) (total int64, err error) {
	if categoryId == 0 {
		// 查询所有电影
		err = dao.DB.Model(&model.Movie{}).Count(&total).Error
	} else {
		err = dao.DB.Model(&model.Movie{}).Where("category_id LIKE ? ", "%"+strconv.Itoa(int(categoryId))+"%").Count(&total).Error
	}
	return
}

func (dao *MovieDao) CountHotMovieByCondition(categoryId uint) (total int64, err error) {
	if categoryId == 0 {
		// 查询所有电影
		err = dao.DB.Model(&model.Movie{}).Count(&total).Error
	} else {
		err = dao.DB.Model(&model.Movie{}).Where("category_id LIKE ? and  on_sale = 1 ", "%"+strconv.Itoa(int(categoryId))+"%").Count(&total).Error
	}
	return
}

func (dao *MovieDao) CountUnreleasedMovieByCondition(categoryId uint) (total int64, err error) {
	if categoryId == 0 {
		// 查询所有电影
		err = dao.DB.Model(&model.Movie{}).Count(&total).Error
	} else {
		err = dao.DB.Model(&model.Movie{}).Where("category_id LIKE ? and  on_sale = 0", "%"+strconv.Itoa(int(categoryId))+"%").Count(&total).Error
	}
	return
}

func (dao *MovieDao) ListMovieByCondition(categoryId uint, page model.BasePage) (movies []*model.Movie, err error) {
	if categoryId == 0 {
		// 查询所有电影
		err = dao.DB.Model(&model.Movie{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	} else {
		err = dao.DB.Model(&model.Movie{}).Where("category_id LIKE ? ", "%"+strconv.Itoa(int(categoryId))+"%").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	}
	return
}

func (dao *MovieDao) ListHotMovieByCondition(categoryId uint, page model.BasePage) (movies []*model.Movie, err error) {
	if categoryId == 0 {
		// 查询所有电影
		err = dao.DB.Model(&model.Movie{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	} else {
		err = dao.DB.Model(&model.Movie{}).Where("category_id LIKE ?	and  on_sale = 1", "%"+strconv.Itoa(int(categoryId))+"%").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	}
	return
}

func (dao *MovieDao) ListUnreleasedMovieByCondition(categoryId uint, page model.BasePage) (movies []*model.Movie, err error) {
	if categoryId == 0 {
		// 查询所有电影
		err = dao.DB.Model(&model.Movie{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	} else {
		err = dao.DB.Model(&model.Movie{}).Where("category_id LIKE ?	and  on_sale = 0", "%"+strconv.Itoa(int(categoryId))+"%").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	}
	return
}

func (dao *MovieDao) ListMovieBySales(page model.BasePage) (movies []*model.Movie, err error) {
	err = dao.DB.Order("sales desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&movies).Error
	return
}

func (dao *MovieDao) SearchMovie(info string, page model.BasePage) (products []*model.Movie, err error) {
	err = dao.DB.Model(&model.Movie{}).
		Where("chinese_name LIKE ? OR introduction LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *MovieDao) GetMovieByMovieID(id uint) (movie *model.Movie, err error) {
	err = dao.DB.Model(&model.Movie{}).Where("id=?", id).First(&movie).Error
	return
}

func (dao *MovieDao) AddMovieSales(id uint, price uint) (err error) {
	err = dao.DB.Model(&model.Movie{}).Where("id=?", id).Update("sales", gorm.Expr("sales + ?", price)).Error
	return
}
