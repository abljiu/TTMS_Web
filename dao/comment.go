package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{NewDBClient(ctx)}
}

func NewCommentDaoByDB(db *gorm.DB) *CommentDao {
	return &CommentDao{db}
}

func (dao *CommentDao) CreateComment(Comment *model.Comment) error {
	return dao.DB.Model(&model.Comment{}).Create(&Comment).Error
}

func (dao *CommentDao) GetCommentByID(id uint) (Comment *model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("id=?", id).First(&Comment).Error
	return
}

func (dao *CommentDao) DeleteCommentByID(id uint) (Comment *model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("id=?", id).Delete(&Comment).Error
	return
}
func (dao *CommentDao) CountComment() (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Count(&total).Error
	return
}

func (dao *CommentDao) ListComment(page model.BasePage, sortBy []string, book int) (Comments []*model.Comment, err error) {
	orderStr := ""
	//book 0 没有额外操作 ；1 当rate>=8；2 当rate<=5
	for _, field := range sortBy {
		if orderStr != "" {
			orderStr += " desc,"
		}
		orderStr += field
	}
	if book == 1 {
		err = dao.DB.Model(&model.Comment{}).
			Where("rate >= 8").
			Order(orderStr).
			Offset((page.PageNum - 1) * page.PageSize).
			Limit(page.PageSize).
			Find(&Comments).
			Error
	} else if book == 2 {
		err = dao.DB.Model(&model.Comment{}).
			Where("rate <= 5").
			Order(orderStr).
			Offset((page.PageNum - 1) * page.PageSize).
			Limit(page.PageSize).
			Find(&Comments).
			Error
	} else {
		err = dao.DB.Model(&model.Comment{}).
			Order(orderStr).
			Offset((page.PageNum - 1) * page.PageSize).
			Limit(page.PageSize).
			Find(&Comments).
			Error
	}
	return
}

func (dao *CommentDao) SearchComment(info string, page model.BasePage) (products []*model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).
		Where("content LIKE ?", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}
