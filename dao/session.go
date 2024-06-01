package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
	"sync"
)

type SessionDao struct {
	*gorm.DB
	mu sync.RWMutex `gorm:"-"`
}

func NewSessionDao(ctx context.Context) *SessionDao {
	return &SessionDao{DB: NewDBClient(ctx)}
}

func NewSessionDaoByDB(db *gorm.DB) *SessionDao {
	return &SessionDao{DB: db}
}

func (dao *SessionDao) AddSession(session *model.Session) error {
	dao.mu.Lock()         // 获取写锁
	defer dao.mu.Unlock() // 操作结束后释放写锁
	return dao.DB.Model(&model.Session{}).Create(&session).Error
}

func (dao *SessionDao) GetSessionByID(id uint) (session *model.Session, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).Where("id=?", id).First(&session).Error
	return
}

func (dao *SessionDao) UpdateSessionByID(uid uint, session *model.Session) error {
	dao.mu.Lock()         // 获取写锁
	defer dao.mu.Unlock() // 操作结束后释放写锁
	err := dao.DB.Model(&model.Session{}).Where("id=?", uid).Updates(&session).Error
	return err
}

func (dao *SessionDao) DeleteSessionByID(uid uint) error {
	err := dao.DB.Where("id=?", uid).Delete(&model.Session{}).Error
	return err
}

func (dao *SessionDao) CountSessionByMovieIDAndDate(theaterID uint, movieID uint, date string, curTime string) (total int64, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and movie_id=? and DATE(show_time) like ? and show_time > ?", theaterID, movieID, date, curTime).
		Count(&total).Error
	return
}

func (dao *SessionDao) CountSessionByDate(theaterID uint, date string, curTime string) (total int64, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).Where("theater_id=?  and DATE(show_time) like ?  and show_time > ?", theaterID, date, curTime).Count(&total).Error
	return
}

func (dao *SessionDao) CountSessionByMovieID(theaterID uint, movieID uint, curTime string) (total int64, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).Where("theater_id=? and movie_id=?  and show_time > ?", theaterID, movieID, curTime).Count(&total).Error
	return
}

func (dao *SessionDao) CountSession(theaterID uint, curTime string) (total int64, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).Where("theater_id=?  and show_time > ?", theaterID, curTime).Count(&total).Error
	return
}

func (dao *SessionDao) ListSessionByDateAndMovieID(theaterID uint, movieID uint, date string, curTime string, page model.BasePage) (products []*model.Session, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and movie_id=? and DATE(show_time) like ?  and show_time > ?", theaterID, movieID, date, curTime).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *SessionDao) ListSessionByDate(theaterID uint, date string, curTime string, page model.BasePage) (products []*model.Session, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and DATE(show_time) like ?  and show_time > ?", theaterID, date, curTime).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *SessionDao) ListSessionByMovieID(theaterID uint, movieID uint, curTime string, page model.BasePage) (products []*model.Session, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=? and movie_id=?  and show_time > ?", theaterID, movieID, curTime).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *SessionDao) ListSession(theaterID uint, curTime string, page model.BasePage) (products []*model.Session, err error) {
	dao.mu.RLock()         // 获取读锁
	defer dao.mu.RUnlock() // 操作结束后释放读锁
	err = dao.DB.Model(&model.Session{}).
		Where("theater_id=?  and show_time > ?", theaterID, curTime).
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Find(&products).Error
	return
}
