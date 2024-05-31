package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Upvote 进行点赞
func Upvote(c *gin.Context) {
	// 创建一个 Upvote 实例
	createProductService := service.Upvote{}

	// 尝试将请求的数据绑定到 createProductService 中
	if err := c.ShouldBind(&createProductService); err == nil {
		// 调用 createProductService 的 Create 方法来创建评论
		res := createProductService.Upvote(c.Request.Context())
		// 返回创建结果
		c.JSON(http.StatusOK, res)
	} else {
		// 如果绑定数据出错，返回错误信息
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateComment", err)
	}
}

// DownVote 取消点赞
func DownVote(c *gin.Context) {
	DownVoteService := service.DownVote{}
	if err := c.ShouldBind(&DownVoteService); err == nil {
		res := DownVoteService.DownVote(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DownVote", err)
	}
}
