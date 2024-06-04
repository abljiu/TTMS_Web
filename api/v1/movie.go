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

// ListHotMovieByTheater 获取热映电影列表
func ListHotMovieByTheater(c *gin.Context) {
	listMovie := service.MovieService{}
	if err := c.ShouldBind(&listMovie); err == nil {
		res := listMovie.ListHotByTheater(c.Request.Context())
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
	listIndexHotMovies := service.MovieService{}
	if err := c.ShouldBind(&listIndexHotMovies); err == nil {
		res := listIndexHotMovies.ListIndexHotMovies(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListIndexHotMovies", err)
	}
}

// DeleteMovie 获取首页热映电影
func DeleteMovie(c *gin.Context) {
	deleteMovie := service.MovieService{}
	if err := c.ShouldBind(&deleteMovie); err == nil {
		res := deleteMovie.Delete(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteMovie", err)
	}
}

// GetMovie 获取电影详细信息
func GetMovie(c *gin.Context) {
	getMovie := service.MovieService{}
	if err := c.ShouldBind(&getMovie); err == nil {
		res := getMovie.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchMovie", err)
	}
}
