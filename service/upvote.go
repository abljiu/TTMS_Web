package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"gorm.io/gorm"
)

type Upvote struct {
	UserId    uint `json:"user_id" form:"user_id"   binding:"required"`
	CommentId uint `json:"comment_id" form:"comment_id"   binding:"required"`
}
type DownVote struct {
	UpvoteId  uint `json:"upvote_id" form:"upvote_id"   binding:"required"`
	UserId    uint `json:"user_id" form:"user_id"   binding:"required"`
	CommentId uint `json:"comment_id" form:"comment_id"   binding:"required"`
}

// Upvote 对评论进行点赞
func (service *Upvote) Upvote(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	Upvote := &model.Upvote{
		Model:     gorm.Model{},
		UserID:    service.UserId,
		CommentID: service.CommentId,
	}

	UpvoteDao := dao.NewUpvoteDao(ctx)
	err = UpvoteDao.CreateUpvote(Upvote)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateUpvote", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	CommentDao := dao.NewCommentDao(ctx)
	Comment, _ := CommentDao.GetCommentByID(service.CommentId)
	Comment.UpvoteNum = Comment.UpvoteNum + 1
	err = CommentDao.UpdateCommentByID(Comment.ID, Comment)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateComment", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("comment add upvote-num", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUpvote(Upvote),
	}
}

// DownVote 取消点赞
func (service *DownVote) DownVote(ctx context.Context) serializer.Response {
	var err error
	code := e.Success
	DownVoteDao := dao.NewUpvoteDao(ctx)
	CommentDao := dao.NewCommentDao(ctx)
	Comment, _ := CommentDao.GetCommentByID(service.CommentId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("GetCommentByID", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//更新点赞数目
	Comment.UpvoteNum = Comment.UpvoteNum - 1
	err = CommentDao.UpdateCommentByID(Comment.ID, Comment)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UpdateComment", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//通过点赞ID删除
	downVote, err := DownVoteDao.DeleteUpvoteByID(service.UpvoteId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("DownVote", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUpvote(downVote),
	}
}
