package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) GetNoticeByType(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
