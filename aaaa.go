package main

//
//import (
//	"TTMS_Web/dao"
//	"TTMS_Web/model"
//	"TTMS_Web/pkg/e"
//	"TTMS_Web/pkg/util"
//	"TTMS_Web/serializer"
//	"TTMS_Web/service"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"sync"
//)
//
//// GetAcclaims 根据评分逆序，评分相同时根据时间逆序
//func GetAcclaims(c *gin.Context) {
//	GetAcclaimsService := service.GetAcclaims{}
//	if err := c.ShouldBind(&GetAcclaimsService); err == nil {
//		res := GetAcclaimsService.GetAcclaims(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		util.LogrusObj.Infoln("GetAcclaims", err)
//	}
//}
//
//// GetAcclaims 获取好评列表
//func (service *GetAcclaims) GetAcclaims(ctx context.Context) serializer.Response {
//	var Comments []*model.Comment
//	var err error
//	code := e.Success
//	if service.PageSize == 0 {
//		service.PageSize = 15
//	}
//	productDao := dao.NewCommentDao(ctx)
//	total, err := productDao.CountComment()
//	if err != nil {
//		code = e.Error
//		util.LogrusObj.Infoln("CountCommentByCondition", err)
//		return serializer.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//		}
//	}
//	wg := new(sync.WaitGroup)
//	wg.Add(1)
//	go func() {
//		productDao = dao.NewCommentDaoByDB(productDao.DB)
//		Comments, _ = productDao.ListComment(service.BasePage, []string{"rate", "created_at"}, GetAcclaim)
//		wg.Done()
//	}()
//	wg.Wait()
//	fmt.Println("success")
//	return serializer.BuildListResponse(serializer.BuildComments(Comments), uint(total))
//}
//
//type GetAcclaims struct {
//	Content   string `json:"content" form:"content"`
//	UserId    uint   `json:"user_id" form:"user_id"  binding:"required"`
//	RlyId     uint   `json:"rly_id" form:"rly_id"`
//	MovieID   uint   `json:"movie_id" form:"movie_id" binding:"required"`
//	Rate      int    `json:"rate" form:"rate"`
//	UpvoteNum int    `json:"upvote_num" form:"upvote_num"`
//	IP        string `json:"ip" form:"ip" binding:"required"`
//	model.BasePage
//}
