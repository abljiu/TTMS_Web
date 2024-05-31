package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type UpvoteDao struct {
	*gorm.DB
}

func NewUpvoteDao(ctx context.Context) *UpvoteDao {
	return &UpvoteDao{NewDBClient(ctx)}
}

func NewUpvoteDaoByDB(db *gorm.DB) *UpvoteDao {
	return &UpvoteDao{db}
}

func (dao *UpvoteDao) CreateUpvote(Upvote *model.Upvote) error {
	return dao.DB.Model(&model.Upvote{}).Create(&Upvote).Error
}

func (dao *UpvoteDao) GetUpvoteByID(id uint) (Upvote *model.Upvote, err error) {
	err = dao.DB.Model(&model.Upvote{}).Where("id=?", id).First(&Upvote).Error
	return
}

func (dao *UpvoteDao) DeleteUpvoteByID(id uint) (Upvote *model.Upvote, err error) {
	err = dao.DB.Model(&model.Upvote{}).Where("id=?", id).Delete(&Upvote).Error
	return
}
func (dao *UpvoteDao) CountUpvote() (total int64, err error) {
	err = dao.DB.Model(&model.Upvote{}).Count(&total).Error
	return
}

func (dao *UpvoteDao) ListUpvote(page model.BasePage) (Upvotes []*model.Upvote, err error) {
	err = dao.DB.Model(&model.Upvote{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&Upvotes).Error
	return
}

// 根据user_id得到comment_id
func (dao *UpvoteDao) GetComments(userID uint) (Upvotes []*model.Upvote, err error) {
	err = dao.DB.Model(&model.Upvote{}).Where("user_id = ?", userID).Find(&Upvotes).Error
	return
}
