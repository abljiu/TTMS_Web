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
	createMovie := service.MovieService{}
	if err := c.ShouldBind(&createMovie); err == nil {
		res := createMovie.Create(c.Request.Context(), movieImg, directorImg, actorImg)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateMovie", err)
	}
}

// ListHotMovie 获取热映电影列表
func ListHotMovie(c *gin.Context) {
	listMovie := service.MovieService{}
	if err := c.ShouldBind(&listMovie); err == nil {
		res := listMovie.ListHot(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovie", err)
	}
}

// ListUnreleasedMovie 获取未上映电影列表
func ListUnreleasedMovie(c *gin.Context) {
	listMovie := service.MovieService{}
	if err := c.ShouldBind(&listMovie); err == nil {
		res := listMovie.ListUnreleased(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovie", err)
	}
}

// ListMovie 获取电影列表
func ListMovie(c *gin.Context) {
	listMovie := service.MovieService{}
	if err := c.ShouldBind(&listMovie); err == nil {
		res := listMovie.ListAll(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovie", err)
	}
}

// ListMovieSales 获取电影票房列表
func ListMovieSales(c *gin.Context) {
	listMovieSales := service.MovieService{}
	if err := c.ShouldBind(&listMovieSales); err == nil {
		res := listMovieSales.ListSales(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListMovieSales", err)
	}
}

// SearchMovie 搜索电影
func SearchMovie(c *gin.Context) {
	searchMovie := service.MovieService{}
	if err := c.ShouldBind(&searchMovie); err == nil {
		res := searchMovie.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchMovie", err)
	}
}

// ListIndexHotMovies 获取首页热映电影
func ListIndexHotMovies(c *gin.Context) {
	searchMovie := service.MovieService{}
	if err := c.ShouldBind(&searchMovie); err == nil {
		res := searchMovie.ListIndexHotMovies(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchMovie", err)
	}
}
