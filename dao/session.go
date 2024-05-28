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
