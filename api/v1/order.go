package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubmitOrder(c *gin.Context) {
	var submitOrder service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&submitOrder); err == nil {
		res := submitOrder.Submit(c.Request.Context(), claims.UserID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SubmitOrder", err)
	}
}

func ConfirmOrder(c *gin.Context) {
	var confirmOrder service.OrderService
	if err := c.ShouldBind(&confirmOrder); err == nil {
		res := confirmOrder.Confirm(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ReturnOrder", err)
	}
}

func ReturnOrder(c *gin.Context) {
	var returnOrder service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&returnOrder); err == nil {
		res := returnOrder.Return(c.Request.Context(), claims.UserID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ReturnOrder", err)
	}
}

func PayOrder(c *gin.Context) {
	var payOrder service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&payOrder); err == nil {
		res := payOrder.Pay(c.Request.Context(), claims.UserID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ReturnOrder", err)
	}
}

func GetOrders(c *gin.Context) {
	var payOrder service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&payOrder); err == nil {
		res := payOrder.Get(c.Request.Context(), claims.UserID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ReturnOrder", err)
	}
}
