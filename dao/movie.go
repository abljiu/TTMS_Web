package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
	"strconv"
	"time"
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

func (dao *MovieDao) SearchMovieExactly(info string) (products *model.Movie, err error) {
	err = dao.DB.Model(&model.Movie{}).
		Where("chinese_name = ? ", info).Find(&products).Error
	return
}

func (dao *MovieDao) SearchMovie(info string, page model.BasePage) (products []*model.Movie, err error) {
	err = dao.DB.Model(&model.Movie{}).
		Where("chinese_name LIKE ? ", "%"+info+"%").
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

func (dao *MovieDao) UpdateMovie(id uint, movie *model.Movie) (err error) {
	err = dao.DB.Model(&model.Movie{}).Where("id=?", id).Updates(movie).Error
	return
}

func (dao *MovieDao) AddMovieTheater(movieTheater *model.MovieTheater) (err error) {
	err = dao.DB.Model(&model.MovieTheater{}).Create(movieTheater).Error
	return
}

func (dao *MovieDao) CountHotMovieByTheater(theaterId uint) (total int64, err error) {
	err = dao.DB.Model(&model.MovieTheater{}).Preload("Movie").Preload("Theater").
		Where("theater_id = ?", theaterId).Count(&total).Error
	return
}

func (dao *MovieDao) ListHotMovieByTheater(theaterId uint) (movies []*model.MovieTheater, err error) {
	err = dao.DB.Model(&model.MovieTheater{}).Preload("Movie").Preload("Theater").
		Where("theater_id = ?", theaterId).Find(&movies).Error
	return
}

//
//func (dao *MovieDao) CountIndexHotMovie(today, preDate string) (total int64, err error) {
//	err = dao.DB.Model(&model.Movie{}).
//		Where("on_sale = ? AND show_time BETWEEN ? AND ?", 1, preDate, today).
//		Count(&total).Error
//	return
//}
//
//func (dao *MovieDao) ListIndexHotMovie(today, preDate string, page int) (movies []*model.Movie, err error) {
//	err = dao.DB.Model(&model.Movie{}).
//		Where("on_sale = ? AND show_time BETWEEN ? AND ?", 1, preDate, today).
//		Order("sales desc").Limit(page).Find(movies).Error
//	return
//}

func (dao *MovieDao) CountIndexHotMovie(today, preDate time.Time) (int64, error) {
	var total int64

	// 注意：这里调用 Err() 方法（如果 GORM 版本支持）或者直接访问 Error 字段
	err := dao.DB.Model(&model.Movie{}).
		Where("on_sale = ? AND show_time BETWEEN ? AND ?", 1, preDate, today).
		Count(&total).Error

	return total, err
}

func (dao *MovieDao) ListIndexHotMovie(today, preDate time.Time, size int) (movies []*model.Movie, err error) {
	err = dao.DB.Model(&model.Movie{}).
		Where("on_sale = ? AND show_time BETWEEN ? AND ?", 1, preDate, today).
		Order("sales desc").Limit(size).Find(&movies).Error
	return
}
