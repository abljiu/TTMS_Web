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
		//获取电影详细信息
		v1.POST("get-movie", api.GetMovie)
		//根据类型查询热映电影
		v1.POST("hot-movies", api.ListHotMovie)
		//根据类型查询未上映电影
		v1.POST("unreleased-movies", api.ListUnreleasedMovie)
		//根据类型查询全部电影
		v1.POST("all-movies", api.ListMovie)
		//查询电影票房
		v1.GET("sales", api.ListMovieSales)
		//获取首页热映电影
		v1.GET("index_hot_movies", api.ListIndexHotMovies)
		//根据sessionID获取某场次详细信息
		v1.GET("session", api.GetSession)
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
			authed.GET("show-user", api.ShowData)
			authed.POST("add-money", api.AddUserMoney)

			//提交订单
			authed.POST("submit-order", api.SubmitOrder)
			//确认订单
			authed.POST("confirm-order", api.ConfirmOrder)
			//支付订单
			authed.POST("pay-order", api.PayOrder)
			//退票
			authed.DELETE("return-order", api.ReturnOrder)
			//用户查看个人订单
			authed.GET("orders", api.GetOrders)

			//搜索电影
			authed.POST("movies", api.SearchMovie)
			//用户发布评论，评分
			authed.POST("publishComment", api.PublishComment)
			//用户点赞某个评论	：为每条评论添加字段，判断是否为该用户点赞
			authed.POST("upvote", api.Upvote)
			//用户取消点赞
			authed.DELETE("downVote", api.DownVote)
			//用户回复评论 ：任何用户都可以回复
			authed.POST("replyComment", api.ReplyComment)
			//用户查看影片所有影评
			authed.GET("getCommentsByMovie", api.GetCommentsByMovie)
			//用户查看影片所有影评，点赞数倒叙
			authed.GET("getCommentsByHeat", api.GetCommentsByHeat)
			//用户查看影片所有影评，好评和时间倒叙
			authed.GET("getAcclaims", api.GetAcclaims)
			//用户查看影片所有影评 :按照时间倒序；点赞数倒叙；好评（差）和时间倒叙
			authed.GET("getNegativeComments", api.GetNegativeComments)
			//用户删除自己的评论
			authed.DELETE("deleteCommentById", api.DeleteCommentByID)

			//管理员权限
			admin := v1.Group("/admin") //api/v1/admin
			admin.Use(middleware.Admin())
			{
				//添加管理员
				admin.PUT("add-admin", api.AddAdmin)
				//添加售票员
				admin.PUT("add-conductor", api.AddConductor)

				//添加电影
				admin.POST("add-movie", api.CreateMovie)
				//删除电影
				admin.DELETE("del-movie", api.DeleteMovie)
				//增加场次
				admin.POST("add-session", api.AddSession)
				//修改场次
				admin.PUT("alter-session", api.AlterSession)
				//删除场次
				admin.DELETE("delete-session", api.DeleteSession)
				//获取某影厅场次列表
				admin.GET("sessions", api.ListSession)
				//根据sessionID获取某场次详细信息
				admin.GET("session", api.GetSession)

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
				//添加剧院
				admin.POST("createTheater", api.CreateTheater)
				//修改剧院
				admin.PUT("updateTheater", api.UpdateTheater)
				//删除剧院
				admin.DELETE("deleteTheater", api.DeleteTheater)
				//查找剧院 根据名称
				admin.GET("searchTheater", api.SearchTheater)
				//查找剧院 根据id
				admin.GET("searchTheaterById", api.SearchTheaterById)
				//管理员查看所有影片评论
				admin.GET("getAllComments", api.GetAllComments)
				//管理员通过评论ID查看评论
				admin.GET("getCommentByID", api.GetCommentByID)
				//管理员查看用户的影片评论
				admin.GET("getCommentsByUserId", api.GetCommentsByUserId)
				//管理员根据ID删除评论
				admin.DELETE("deleteCommentById", api.DeleteCommentByID)
				//管理员删除不合法内容的评论
				admin.DELETE("deleteCommentsByContent", api.DeleteCommentsByContent)

			}
		}

	}

	return r
}
