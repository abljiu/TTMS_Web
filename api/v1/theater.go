package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTheater 创建剧院
func CreateTheater(c *gin.Context) {
	// 创建一个 TheaterService 实例
	createProductService := service.TheaterService{}

	// 尝试将请求的数据绑定到 createProductService 中
	if err := c.ShouldBind(&createProductService); err == nil {
		// 调用 createProductService 的 Create 方法来创建剧院
		res := createProductService.Create(c.Request.Context())
		// 返回创建结果
		c.JSON(http.StatusOK, res)
	} else {
		// 如果绑定数据出错，返回错误信息
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateTheater", err)
	}
}

// ListTheater 获取剧院列表
func ListTheater(c *gin.Context) {
	listTheaterService := service.TheaterService{}
	if err := c.ShouldBind(&listTheaterService); err == nil {
		res := listTheaterService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListTheater", err)
	}
}

// SearchTheater 获取剧院列表
func SearchTheater(c *gin.Context) {
	SearchTheaterService := service.TheaterService{}
	if err := c.ShouldBind(&SearchTheaterService); err == nil {
		fmt.Println(SearchTheaterService)
		res := SearchTheaterService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchTheater", err)
	}
}

// SearchTheater 获取剧院列表通过id
func SearchTheaterById(c *gin.Context) {
	SearchTheaterService := service.TheaterService{}
	if err := c.ShouldBind(&SearchTheaterService); err == nil {
		fmt.Println(SearchTheaterService)
		res := SearchTheaterService.SearchById(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchTheater", err)
	}
}

// UpdateTheater 更新剧院
func UpdateTheater(c *gin.Context) {
	UpdateTheaterService := service.TheaterService{}
	if err := c.ShouldBind(&UpdateTheaterService); err == nil {
		res := UpdateTheaterService.Update(c.Request.Context(), UpdateTheaterService.TheaterId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("UpdateTheater", err)
	}
}

// DeleteTheater 删除剧院
func DeleteTheater(c *gin.Context) {
	DeleteTheaterService := service.TheaterService{}
	if err := c.ShouldBind(&DeleteTheaterService); err == nil {
		res := DeleteTheaterService.Delete(c.Request.Context(), DeleteTheaterService.TheaterId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteTheater", err)
	}
}
