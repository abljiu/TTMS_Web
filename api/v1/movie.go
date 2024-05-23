package v1

import (
	"TTMS_Web/pkg/util"
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateMovie 创建电影
func CreateMovie(c *gin.Context) {
	form, _ := c.MultipartForm()
	movieImg := form.File["movie_img"]
	directorImg := form.File["director_img"]
	actorImg := form.File["actor_img"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.MovieService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.UserID, movieImg, directorImg, actorImg)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateMovie", err)
	}
}

// ListMovie 获取电影列表
func ListMovie(c *gin.Context) {
	listMovieService := service.MovieService{}
	if err := c.ShouldBind(&listMovieService); err == nil {
		res := listMovieService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovie", err)
	}
}

// SearchMovie 搜索电影
func SearchMovie(c *gin.Context) {
	SearchMovieService := service.MovieService{}
	if err := c.ShouldBind(&SearchMovieService); err == nil {
		res := SearchMovieService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovie", err)
	}
}
