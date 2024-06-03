package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

const (
	NoFunc       = 0
	GetAcclaim   = 1
	GetUnAcclaim = 2
	GetByUser    = 3
	GetByMovie   = 4
)

// 根据内容搜索全部评论
type SearchComment struct {
	Content string `json:"content" form:"content" binding:"required"`
	model.BasePage
}

// 发布评论
type PublishComment struct {
	CommentId uint   `json:"comment_id" form:"comment_id"`
	Content   string `json:"content" form:"content"`
	UserId    uint   `json:"user_id" form:"user_id" binding:"required"`
	MovieID   uint   `json:"movie_id" form:"movie_id" binding:"required"`
	Rate      int    `json:"rate" form:"rate"`
	IP        string `json:"ip" form:"ip" binding:"required"`
	model.BasePage
}

// 回复评论
type ReplyComment struct {
	CommentId uint   `json:"comment_id" form:"comment_id"`
	Content   string `json:"content" form:"content"  binding:"required"`
	UserId    uint   `json:"user_id" form:"user_id" binding:"required"`
	RlyId     uint   `json:"rly_id" form:"rly_id"  binding:"required"`
	MovieID   uint   `json:"movie_id" form:"movie_id"`
	IP        string `json:"ip" form:"ip" binding:"required"`
	model.BasePage
}

// 得到影片的全部评论
type GetCommentsByMovie struct {
	MovieID uint `json:"movie_id" form:"movie_id" binding:"required"`
	UserId  uint `json:"user_id" form:"user_id" binding:"required"`
	model.BasePage
}

// 通过ID得到评论
type GetCommentByID struct {
	CommentId uint `json:"comment_id" form:"comment_id" binding:"required"`
	model.BasePage
}
type GetCommentsByHeat struct {
	UserId  uint `json:"user_id" form:"user_id"  binding:"required"`
	MovieID uint `json:"movie_id" form:"movie_id" binding:"required"`
	model.BasePage
}
type GetAcclaims struct {
	UserId  uint `json:"user_id" form:"user_id"  binding:"required"`
	MovieID uint `json:"movie_id" form:"movie_id" binding:"required"`
	model.BasePage
}
type GetNegativeComments struct {
	UserId  uint `json:"user_id" form:"user_id"  binding:"required"`
	MovieID uint `json:"movie_id" form:"movie_id" binding:"required"`
	model.BasePage
}
type GetAllComments struct {
	model.BasePage
}
type GetCommentsByUserId struct {
	UserId uint `json:"user_id" form:"user_id"  binding:"required"`
	model.BasePage
}
type DeleteCommentByID struct {
	CommentId uint `json:"comment_id" form:"comment_id"  binding:"required"`
	UserId    uint `json:"user_id" form:"user_id"  binding:"required"`
	MovieID   uint `json:"movie_id" form:"movie_id" binding:"required"`
	model.BasePage
}
type DeleteCommentsByContent struct {
	Content string `json:"content" form:"content"  binding:"required"`
	model.BasePage
}

// PublishComment 发布评论
func (service *PublishComment) PublishComment(ctx context.Context) serializer.Response {
	var err error
	code := e.Success
	//判断用户是否看过该影片?
	//productDao:=dao.NewOrderDao(ctx)
	//productDao.
	Comment := &model.Comment{
		Model:   gorm.Model{},
		Content: service.Content,
		UserID:  service.UserId,
		MovieID: service.MovieID,
		Rate:    uint(service.Rate),
		IP:      service.IP,
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

// ReplyComment 回复评论
func (service *ReplyComment) ReplyComment(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	Comment := &model.Comment{
		Model:   gorm.Model{},
		Content: service.Content,
		UserID:  service.UserId,
		RlyID: sql.NullInt64{
			Int64: int64(service.RlyId),
			Valid: true,
		},
		MovieID: service.MovieID,
		IP:      service.IP,
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

// GetAcclaims 获取好评列表，显示自己点赞的评论
func (service *GetAcclaims) GetAcclaims(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountAcclaims()
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
		Comments, _ = productDao.ListComment(service.BasePage, []string{"rate", "created_at"}, []uint{0, 0, 0, service.MovieID, 0}, GetAcclaim)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return JudgeUpvoteIsSelf(ctx, service.UserId, total, Comments)
}

// JudgeUpvoteIsSelf 根据userID得到自己点赞过的评论
func JudgeUpvoteIsSelf(ctx context.Context, userID uint, total int64, comments []*model.Comment) serializer.Response {
	//根据user_id得到comment_id
	productDao := dao.NewUpvoteDao(ctx)
	upvotes, _ := productDao.GetComments(userID)
	//将comment字段进行标记
	for _, i := range upvotes {
		for _, j := range comments {
			if j.ID == i.CommentID {
				j.IsSelfUpvote = true
			}
		}
	}
	fmt.Println("______comments")
	fmt.Println(serializer.BuildComments(comments)[0].IsSelfUpvote)
	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}

// GetNegativeComments 获取差评列表，并按照时间倒序
func (service *GetNegativeComments) GetNegativeComments(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountBadComments()
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
		Comments, _ = productDao.ListComment(service.BasePage, []string{"rate", "created_at"}, []uint{0, 0, 0, service.MovieID, 0}, GetUnAcclaim)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return JudgeUpvoteIsSelf(ctx, service.UserId, total, Comments)
}

// List 获取评论列表
func (service *GetAllComments) List(ctx context.Context) serializer.Response {
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
		Comments, _ = productDao.ListComment(service.BasePage, []string{"created_at"}, []uint{0, 0, 0, 0, 0}, NoFunc)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}

// GetCommentsByMovie 根据影片得到评论
func (service *GetCommentsByMovie) GetCommentsByMovie(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountCommentByMovieID(service.MovieID)
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
		Comments, _ = productDao.ListComment(service.BasePage, []string{"created_at"}, []uint{0, 0, 0, service.MovieID, 0}, GetByMovie)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return JudgeUpvoteIsSelf(ctx, service.UserId, total, Comments)
}

// GetCommentsByHeat 根据热度得到评论
func (service *GetCommentsByHeat) GetCommentsByHeat(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountCommentByMovieID(service.MovieID)
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
		Comments, _ = productDao.ListComment(service.BasePage, []string{"upvote_num", "created_at"}, []uint{0, 0, 0, service.MovieID, 0}, GetByMovie)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return JudgeUpvoteIsSelf(ctx, service.UserId, total, Comments)
}

// GetCommentsByUserId 根据用户ID获取好评列表
func (service *GetCommentsByUserId) GetCommentsByUserId(ctx context.Context) serializer.Response {
	var Comments []*model.Comment
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewCommentDao(ctx)
	total, err := productDao.CountCommentByUserID(service.UserId)
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
		Comments, _ = productDao.ListComment(service.BasePage, []string{"created_at"}, []uint{0, 0, service.UserId, 0, 0}, GetByUser)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
}

// Search 根据内容搜索评论
func (service *SearchComment) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewCommentDao(ctx)
	//查找不法词汇，其他所需词汇
	Comments, err := productDao.SearchComment(service.Content, service.BasePage)
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

// GetCommentByID 通过ID得到评论
func (service *GetCommentByID) GetCommentByID(ctx context.Context) serializer.Response {
	code := e.Success

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

// DeleteCommentByID 根据id删除信息
func (service *DeleteCommentByID) DeleteCommentByID(ctx context.Context, uid uint) serializer.Response {
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

// DeleteCommentsByContent 根据内容删除信息
func (service *DeleteCommentsByContent) DeleteCommentsByContent(ctx context.Context) serializer.Response {
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	CommentDao := dao.NewCommentDao(ctx)
	Comments, err := CommentDao.DeleteCommentsByContent(service.Content, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(len(Comments)))
}
