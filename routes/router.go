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

		//查询电影票房
		v1.GET("sales", api.ListMovieSales)
		//根据类型查询热映电影
		v1.GET("hot-movies", api.ListHotMovie)
		//根据类型查询未上映电影
		v1.GET("unreleased-movies", api.ListUnreleasedMovie)
		//根据类型查询全部电影
		v1.GET("all-movies", api.ListMovie)
		//根据sessionId返回场次信息
		v1.GET("session", api.GetSession)

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
			//authed.POST("pay-order",api.)
			authed.DELETE("return-order", api.ReturnOrder)
			authed.GET("confirm-order", api.ConfirmOrder)

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
				//获取影院热映电影列表
				admin.GET("movie/getHot", api.ListHotMovieByTheater)

				//某剧院的影厅列表
				admin.GET("halls", api.ListHall)
				//创建影厅
				admin.POST("hall/create", api.CreateHall)
				//删除影厅 根据影厅id
				admin.DELETE("hall/delete", api.DeleteHall)
				//更新影厅信息
				admin.PUT("hall/update", api.UpdateHall)
				//影厅详细信息
				admin.GET("hall", api.GetHall)
			}
		}

	}

	return r
}
