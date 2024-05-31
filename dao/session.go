package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type SessionDao struct {
	*gorm.DB
}

func NewSessionDao(ctx context.Context) *SessionDao {
	return &SessionDao{NewDBClient(ctx)}
}

func NewSessionDaoByDB(db *gorm.DB) *SessionDao {
	return &SessionDao{db}
}

func (dao *SessionDao) AddSession(session *model.Session) error {
	return dao.DB.Model(&model.Session{}).Create(&session).Error
}

func (dao *SessionDao) GetSessionByID(id uint) (session *model.Session, err error) {
	err = dao.DB.Model(&model.Session{}).Where("id=?", id).First(&session).Error
	return
}

func (dao *SessionDao) UpdateSessionByID(uid uint, session *model.Session) error {
	err := dao.DB.Model(&model.Session{}).Where("id=?", uid).Updates(&session).Error
	return err
}

func (dao *SessionDao) DeleteSessionByID(uid uint) error {
	err := dao.DB.Where("id=?", uid).Delete(&model.Session{}).Error
	return err
}

func (dao *SessionDao) CountSessionByMovieIDAndDate(theaterID uint, movieID uint, date string) (total int64, err error) {
	err = dao.DB.Model(&model.Session{}).Where("theater_id=? and movie_id=? and show_time like ?", theaterID, movieID, date).Count(&total).Error
	return
}

func (dao *SessionDao) CountSessionByDate(theaterID uint, date string) (total int64, err error) {
	err = dao.DB.Model(&model.Session{}).Where("theater_id=?  and show_time like ?", theaterID, date).Count(&total).Error
	return
}

func (dao *SessionDao) CountSessionByMovieID(theaterID uint, movieID uint) (total int64, err error) {
	err = dao.DB.Model(&model.Session{}).Where("theater_id=? and movie_id=?", theaterID, movieID).Count(&total).Error
	return
}

func (dao *SessionDao) CountSession(theaterID uint) (total int64, err error) {
	err = dao.DB.Model(&model.Session{}).Where("theater_id=? ", theaterID).Count(&total).Error
	return
}

func (dao *SessionDao) ListSessionByDateAndMovieID(theaterID uint, movieID uint, date string, page model.BasePage) (products []*model.Session, err error) {
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and movie_id=? and DATE(show_time) like ?", theaterID, movieID, date).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *SessionDao) ListSessionByDate(theaterID uint, date string, page model.BasePage) (products []*model.Session, err error) {
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and DATE(show_time) like ?", theaterID, date).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *SessionDao) ListSessionByMovieID(theaterID uint, movieID uint, page model.BasePage) (products []*model.Session, err error) {
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and movie_id=? ", theaterID, movieID).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *SessionDao) ListSession(theaterID uint, page model.BasePage) (products []*model.Session, err error) {
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? ", theaterID).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}
