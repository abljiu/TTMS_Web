package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddSession 添加场次
func AddSession(c *gin.Context) {
	addSession := service.SessionServer{}
	if err := c.ShouldBind(&addSession); err == nil {
		res := addSession.Add(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("AddSession", err)
	}
}

// AlterSession 修改场次
func AlterSession(c *gin.Context) {
	alterSession := service.SessionServer{}
	if err := c.ShouldBind(&alterSession); err == nil {
		res := alterSession.Alter(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("AlterSession", err)
	}
}

// DeleteSession 删除场次
func DeleteSession(c *gin.Context) {
	deleteSession := service.SessionServer{}
	if err := c.ShouldBind(&deleteSession); err == nil {
		res := deleteSession.Delete(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteSession", err)
	}
}

// ListSession 获取某影厅场次列表
func ListSession(c *gin.Context) {
	listSession := service.SessionListRequest{}
	if err := c.ShouldBind(&listSession); err == nil {
		res := listSession.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteSession", err)
	}
}

// GetSession 获取某场次详细信息
func GetSession(c *gin.Context) {
	getSession := service.SessionRequest{}
	if err := c.ShouldBind(&getSession); err == nil {
		res := getSession.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteSession", err)
	}
}
