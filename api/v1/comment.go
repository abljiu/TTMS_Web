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

// ReplyComment 发布评论
func ReplyComment(c *gin.Context) {
	// 创建一个 ReplyComment 实例
	createProductService := service.ReplyComment{}
	// 尝试将请求的数据绑定到 createProductService 中
	if err := c.ShouldBind(&createProductService); err == nil {
		// 调用 createProductService 的 Create 方法来创建评论
		res := createProductService.ReplyComment(c.Request.Context())
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

// GetNegativeComments 根据评分逆序，评分相同时根据时间逆序
func GetNegativeComments(c *gin.Context) {
	GetNegativeCommentsService := service.GetNegativeComments{}
	if err := c.ShouldBind(&GetNegativeCommentsService); err == nil {
		res := GetNegativeCommentsService.GetNegativeComments(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetNegativeComments", err)
	}
}

// GetAllComments 获取所有的评论
func GetAllComments(c *gin.Context) {
	GetAllCommentsService := service.GetAllComments{}
	if err := c.ShouldBind(&GetAllCommentsService); err == nil {
		res := GetAllCommentsService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetAllComments", err)
	}
}

// GetCommentsByUserId 获取所有的评论
func GetCommentsByUserId(c *gin.Context) {
	GetCommentsByUserIdService := service.GetCommentsByUserId{}
	if err := c.ShouldBind(&GetCommentsByUserIdService); err == nil {
		res := GetCommentsByUserIdService.GetCommentsByUserId(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetCommentsByUserId", err)
	}
}

// GetCommentByID 通过id获取评论
func GetCommentByID(c *gin.Context) {
	GetCommentByIDService := service.GetCommentByID{}
	if err := c.ShouldBind(&GetCommentByIDService); err == nil {
		fmt.Println(GetCommentByIDService)
		res := GetCommentByIDService.GetCommentByID(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchTheater", err)
	}
}

// DeleteCommentByID 根据删除评论
func DeleteCommentByID(c *gin.Context) {
	DeleteCommentByIDService := service.DeleteCommentByID{}
	if err := c.ShouldBind(&DeleteCommentByIDService); err == nil {
		res := DeleteCommentByIDService.DeleteCommentByID(c.Request.Context(), DeleteCommentByIDService.CommentId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteCommentByID", err)
	}
}

// DeleteCommentsByContent 根据内容删除评论
func DeleteCommentsByContent(c *gin.Context) {
	DeleteCommentsByContentService := service.DeleteCommentsByContent{}
	if err := c.ShouldBind(&DeleteCommentsByContentService); err == nil {
		res := DeleteCommentsByContentService.DeleteCommentsByContent(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteCommentsByContent", err)
	}
}

// SearchComment 根据内容搜索评论
func SearchComment(c *gin.Context) {
	SearchCommentService := service.SearchComment{}
	if err := c.ShouldBind(&SearchCommentService); err == nil {
		fmt.Println(SearchCommentService)
		res := SearchCommentService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchComment", err)
	}
}
