package routes

import (
	api "TTMS_Web/api/v1"
	"TTMS_Web/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)
		//根据类型查询电影
		v1.GET("movies", api.ListMovie)
		//查询电影票房
		v1.GET("sales", api.ListMovieSales)

		//需要登录保护
		authed := v1.Group("/") //api/v1/
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			authed.POST("submit-order", api.SubmitOrder)
			authed.DELETE("cancel-order", api.CancelOrder)
			authed.DELETE("return-order", api.ReturnOrder)

			//显示金额
			//authed.POST("money", api.ShowMoney)

			//搜索电影
			//authed.POST("movies", api.SearchMovie)

			//管理员权限
			admin := v1.Group("/admin") //api/v1/admin
			admin.Use(middleware.Admin())
			{
				//添加电影
				admin.POST("add-movie", api.CreateMovie)
				//增加场次
				admin.POST("add-session", api.AddSession)
				//修改场次
				admin.PUT("alter-session", api.AlterSession)
				//删除场次
				admin.DELETE("delete-session", api.DeleteSession)
				//添加剧院
				//admin.POST("add-theater", api.AddTheater)
			}
		}
	}
	return r
}
