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
func (dao *CommentDao) UpdateCommentByID(uid uint, Comment *model.Comment) error {
	err := dao.DB.Model(&model.Comment{}).Where("id=?", uid).Updates(&Comment).Error
	return err
}
func (dao *CommentDao) GetCommentByID(id uint) (Comment *model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("id=?", id).First(&Comment).Error
	return
}

func (dao *CommentDao) DeleteCommentByID(id uint) (Comment *model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("id=?", id).Delete(&Comment).Error
	return
}
func (dao *CommentDao) DeleteCommentsByIDs(products []*model.Comment) (Comments []*model.Comment, err error) {
	for _, i := range products {
		comment, err := dao.DeleteCommentByID(i.ID)
		if err == nil {
			return nil, err
		}
		Comments = append(Comments, comment)
	}
	return
}
func (dao *CommentDao) CountComment() (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Count(&total).Error
	return
}
func (dao *CommentDao) CountCommentByUserID(userID uint) (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("user_id = ?", userID).Count(&total).Error
	return
}
func (dao *CommentDao) CountCommentByMovieID(movieID uint) (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("movie_id = ?", movieID).Count(&total).Error
	return
}
func (dao *CommentDao) CountAcclaims() (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("rate >= 8").Count(&total).Error
	return
}
func (dao *CommentDao) CountBadComments() (total int64, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("rate <= 5").Count(&total).Error
	return
}

func (dao *CommentDao) ListComment(page model.BasePage, sortBy []string, ids []uint, book int) (Comments []*model.Comment, err error) {
	orderStr := ""
	//book 0 没有额外操作 ；1 当rate>=8；2 当rate<=5 ；3 user_id=?	4 movie_id=?
	for _, field := range sortBy {
		if orderStr != "" {
			orderStr += " desc,"
		}
		orderStr += field
	}
	if book == 1 {
		err = dao.DB.Model(&model.Comment{}).
			Where("rate >= 8 and movie_id = ?", ids[3]).
			Order(orderStr).
			Offset((page.PageNum - 1) * page.PageSize).
			Limit(page.PageSize).
			Find(&Comments).
			Error
	} else if book == 2 {
		err = dao.DB.Model(&model.Comment{}).
			Where("rate <= 5 and movie_id = ?", ids[3]).
			Order(orderStr).
			Offset((page.PageNum - 1) * page.PageSize).
			Limit(page.PageSize).
			Find(&Comments).
			Error
	} else if book == 3 {
		err = dao.DB.Model(&model.Comment{}).
			Where("user_id = ?", ids[book-1]).
			Order(orderStr).
			Offset((page.PageNum - 1) * page.PageSize).
			Limit(page.PageSize).
			Find(&Comments).
			Error
	} else if book == 4 {
		err = dao.DB.Model(&model.Comment{}).
			Where("movie_id = ?", ids[book-1]).
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
func (dao *CommentDao) DeleteCommentsByContent(info string, page model.BasePage) (products []*model.Comment, err error) {
	products, err = dao.SearchComment(info, page)
	_, err = dao.DeleteCommentsByIDs(products)
	return
}
