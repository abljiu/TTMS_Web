package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListHall 获取影厅列表
func ListHall(c *gin.Context) {
	listHallService := service.HallListRequest{}
	if err := c.ShouldBind(&listHallService); err == nil {
		res := listHallService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovie", err)
	}
}

// CreateHall 创建影厅
func CreateHall(c *gin.Context) {
	createProductService := service.HallCreateRequest{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateMovie", err)
	}
}

// DeleteHall 删除影厅
func DeleteHall(c *gin.Context) {
	deleteProductService := service.HallIDRequest{}
	if err := c.ShouldBind(&deleteProductService); err == nil {
		res := deleteProductService.Delete(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteMovie", err)
	}
}

// UpdateHall 更新影厅信息
func UpdateHall(c *gin.Context) {
	deleteProductService := service.HallUpdateRequest{}
	if err := c.ShouldBind(&deleteProductService); err == nil {
		res := deleteProductService.Update(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteMovie", err)
	}
}

// GetHall 获取某影厅详细信息
func GetHall(c *gin.Context) {
	deleteProductService := service.HallIDRequest{}
	if err := c.ShouldBind(&deleteProductService); err == nil {
		res := deleteProductService.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteMovie", err)
	}
}

//func Mo(c *gin.Context) {
//	deleteProductService := service.ModifySeatService{}
//	if err := c.ShouldBind(&deleteProductService); err == nil {
//		deleteProductService.Modify(c.Request.Context())
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		util.LogrusObj.Infoln("DeleteMovie", err)
//	}
//}
