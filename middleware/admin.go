package middleware

import (
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Admin 验证管理员权限
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if claims.Status != model.Administrator {
			code = e.ErrorStatus
		}
		if code != e.Success {
			c.JSON(http.StatusOK, gin.H{
				"status ": code,
				"msg":     e.GetMsg(code),
			})
			c.Abort()
			return
		}
	}
}
