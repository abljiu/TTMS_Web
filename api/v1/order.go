package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubmitOrder(c *gin.Context) {
	var submitOrder service.OrderService
	if err := c.ShouldBind(&submitOrder); err == nil {
		res := submitOrder.Submit(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SubmitOrder", err)
	}
}

func CancelOrder(c *gin.Context) {
	var cancelOrder service.OrderService
	if err := c.ShouldBind(&cancelOrder); err == nil {
		res := cancelOrder.Cancel(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CancelOrder", err)
	}
}

func ReturnOrder(c *gin.Context) {
	var returnOrder service.OrderService
	if err := c.ShouldBind(&returnOrder); err == nil {
		res := returnOrder.Return(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ReturnOrder", err)
	}
}
