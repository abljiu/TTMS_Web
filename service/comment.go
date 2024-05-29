package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

const (
	NoFunc       = 0
	GetAcclaim   = 1
	GetUnAcclaim = 2
	GetByMovie   = 5
	GetByHeat    = 6
)

type CommentService struct {
	CommentId uint   `json:"comment_id" form:"comment_id"`
	Content   string `json:"content" form:"content" `
	UserId    uint   `json:"user_id" form:"user_id"`
	RlyId     uint   `json:"rly_id" form:"rly_id"`
	MovieID   uint   `json:"movie_id" form:"movie_id"`
	Rate      int    `json:"rate" form:"rate"`
	UpvoteNum int    `json:"upvote_num" form:"upvote_num"`
	IP        string `json:"ip" form:"ip"`
	model.BasePage
}
type PublishComment struct {
	Content   string `json:"content" form:"content"`
	UserId    uint   `json:"user_id" form:"user_id" binding:"required"`
	RlyId     uint   `json:"rly_id" form:"rly_id"`
	MovieID   uint   `json:"movie_id" form:"movie_id" binding:"required"`
	Rate      int    `json:"rate" form:"rate"`
	UpvoteNum int    `json:"upvote_num" form:"upvote_num"`
	IP        string `json:"ip" form:"ip" binding:"required"`
	model.BasePage
}
type GetCommentsByMovie struct {
	Content   string `json:"content" form:"content"`
	UserId    uint   `json:"user_id" form:"user_id"`
	RlyId     uint   `json:"rly_id" form:"rly_id"`
	MovieID   uint   `json:"movie_id" form:"movie_id" binding:"required"`
	Rate      int    `json:"rate" form:"rate"`
	UpvoteNum int    `json:"upvote_num" form:"upvote_num"`
	IP        string `json:"ip" form:"ip" binding:"required"`
	model.BasePage
}
type GetCommentsByHeat struct {
	Content   string `json:"content" form:"content"`
	UserId    uint   `json:"user_id" form:"user_id"  binding:"required"`
	RlyId     uint   `json:"rly_id" form:"rly_id"`
	MovieID   uint   `json:"movie_id" form:"movie_id" binding:"required"`
	Rate      int    `json:"rate" form:"rate"`
	UpvoteNum int    `json:"upvote_num" form:"upvote_num"`
	IP        string `json:"ip" form:"ip" binding:"required"`
	model.BasePage
}
type GetAcclaims struct {
	Content   string `json:"content" form:"content"`
	UserId    uint   `json:"user_id" form:"user_id"  binding:"required"`
	RlyId     uint   `json:"rly_id" form:"rly_id"`
	MovieID   uint   `json:"movie_id" form:"movie_id" binding:"required"`
	Rate      int    `json:"rate" form:"rate"`
	UpvoteNum int    `json:"upvote_num" form:"upvote_num"`
	IP        string `json:"ip" form:"ip" binding:"required"`
	model.BasePage
}

// Create 上传新剧院
func (service *PublishComment) PublishComment(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	Comment := &model.Comment{
		Model: gorm.Model{},
		//Name:    service.CommentName,
		//Address: service.Address,
		//HallNum: service.HallNum,
	}

	CommentDao := dao.NewCommentDao(ctx)
	err = CommentDao.CreateComment(Comment)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateComment", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildComment(Comment),
	}
}

// GetAcclaims 获取好评列表
func (service *GetAcclaims) GetAcclaims(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountComment()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountCommentByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewCommentDaoByDB(productDao.DB)
		Comments, _ = productDao.ListComment(service.BasePage, []string{"rate", "created_at"}, GetAcclaim)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}

// List 获取剧院列表
func (service *CommentService) List(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountComment()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountCommentByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewCommentDaoByDB(productDao.DB)
		Comments, _ = productDao.ListComment(service.BasePage, []string{}, 0)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}

// List 获取剧院列表
func (service *GetCommentsByMovie) GetCommentsByMovie(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountComment()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountCommentByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewCommentDaoByDB(productDao.DB)
		Comments, _ = productDao.ListComment(service.BasePage, []string{}, 0)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}

// List 获取剧院列表
func (service *GetCommentsByHeat) GetCommentsByHeat(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountComment()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountCommentByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewCommentDaoByDB(productDao.DB)
		Comments, _ = productDao.ListComment(service.BasePage, []string{}, 0)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}

// Search 搜索剧院根据名称
func (service *CommentService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewCommentDao(ctx)
	//?
	Comments, err := productDao.SearchComment("", service.BasePage)
	if err != nil {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(len(Comments)))
}

// Search 搜索剧院根据id
func (service *CommentService) SearchById(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewCommentDao(ctx)
	Comment, err := productDao.GetCommentByID(service.CommentId)
	if err != nil {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildComment(Comment),
	}
}

// Delete 删除信息
func (service *CommentService) Delete(ctx context.Context, uid uint) serializer.Response {
	var Comment *model.Comment
	var err error
	code := e.Success
	CommentDao := dao.NewCommentDao(ctx)
	Comment, err = CommentDao.GetCommentByID(uid)
	_, err = CommentDao.DeleteCommentByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildComment(Comment),
	}
}
