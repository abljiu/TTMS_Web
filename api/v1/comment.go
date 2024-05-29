package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PublishComment 发布评论
func PublishComment(c *gin.Context) {
	// 创建一个 PublishComment 实例
	createProductService := service.PublishComment{}

	// 尝试将请求的数据绑定到 createProductService 中
	if err := c.ShouldBind(&createProductService); err == nil {
		// 调用 createProductService 的 Create 方法来创建评论
		res := createProductService.PublishComment(c.Request.Context())
		// 返回创建结果
		c.JSON(http.StatusOK, res)
	} else {
		// 如果绑定数据出错，返回错误信息
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateComment", err)
	}
}

// GetCommentsByMovie 根据movieID获取评论
func GetCommentsByMovie(c *gin.Context) {
	GetCommentsByMovieService := service.GetCommentsByMovie{}
	if err := c.ShouldBind(&GetCommentsByMovieService); err == nil {
		res := GetCommentsByMovieService.GetCommentsByMovie(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetCommentsByMovie", err)
	}
}

// GetCommentsByHeat 根据热度排序
func GetCommentsByHeat(c *gin.Context) {
	GetCommentsByHeatService := service.GetCommentsByHeat{}
	if err := c.ShouldBind(&GetCommentsByHeatService); err == nil {
		res := GetCommentsByHeatService.GetCommentsByHeat(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetCommentsByHeat", err)
	}
}

// GetAcclaims 根据评分逆序，评分相同时根据时间逆序
func GetAcclaims(c *gin.Context) {
	GetAcclaimsService := service.GetAcclaims{}
	if err := c.ShouldBind(&GetAcclaimsService); err == nil {
		res := GetAcclaimsService.GetAcclaims(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetAcclaims", err)
	}
}

// SearchComment 获取评论列表
func SearchComment(c *gin.Context) {
	SearchCommentService := service.CommentService{}
	if err := c.ShouldBind(&SearchCommentService); err == nil {
		fmt.Println(SearchCommentService)
		res := SearchCommentService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchComment", err)
	}
}

// SearchComment 获取评论列表通过id
func SearchCommentById(c *gin.Context) {
	SearchCommentService := service.CommentService{}
	if err := c.ShouldBind(&SearchCommentService); err == nil {
		fmt.Println(SearchCommentService)
		res := SearchCommentService.SearchById(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchComment", err)
	}
}

//// UpdateComment 更新评论
//func UpdateComment(c *gin.Context) {
//	UpdateCommentService := service.CommentService{}
//	if err := c.ShouldBind(&UpdateCommentService); err == nil {
//		res := UpdateCommentService.Update(c.Request.Context(), UpdateCommentService.CommentId)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		util.LogrusObj.Infoln("UpdateComment", err)
//	}
//}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	DeleteCommentService := service.CommentService{}
	if err := c.ShouldBind(&DeleteCommentService); err == nil {
		res := DeleteCommentService.Delete(c.Request.Context(), DeleteCommentService.CommentId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteComment", err)
	}
}
