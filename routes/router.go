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
		//查询电影
		v1.GET("movies", api.ListMovie)
		//查询电影票房
		v1.GET("sales", api.ListMovieSales)
		//获取剧院列表
		v1.GET("theaters", api.ListTheater)
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
			authed.POST("cancel-order", api.CancelOrder)
			authed.POST("return-order", api.ReturnOrder)

			//显示金额
			//authed.POST("money", api.ShowMoney)

			//搜索电影
			//authed.POST("movies", api.SearchMovie)

			//管理员权限
			admin := v1.Group("/admin") //api/v1/admin
			admin.Use(middleware.Admin())
			{
				//添加电影
				admin.POST("movie", api.CreateMovie)
				//添加剧院
				admin.POST("theater", api.CreateTheater)
				//修改剧院
				admin.PUT("theater", api.UpdateTheater)
				//删除剧院
				admin.DELETE("theater", api.DeleteTheater)
				//查找剧院 根据名称
				admin.GET("theater", api.SearchTheater)
				//查找剧院 根据id
				v1.GET("theaterById", api.SearchTheaterById)
			}
		}
	}
	return r
}
