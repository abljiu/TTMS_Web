package service

import (
	"TTMS_Web/conf"
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mime/multipart"
	"strings"
	"sync"
	"time"
)

type MovieService struct {
	MovieId      uint          `json:"movie_id" form:"movie_id"`
	ChineseName  string        `json:"chinese_name" form:"chinese_name"`
	EnglishName  string        `json:"english_name" form:"english_name"`
	CategoryId   []uint        `json:"category_id" form:"category_id"`
	CategoryIds  string        `json:"category_ids" form:"category_ids"`
	Area         string        `json:"area" form:"area"`
	Duration     time.Duration `json:"duration" form:"duration"`
	ShowTime     time.Time     `json:"show_time" form:"show_time" time_format:"2006-01-02"`
	Introduction string        `json:"introduction" form:"introduction"`
	OnSale       bool          `json:"on_sale" form:"on_sale"`
	Score        float64       `json:"score" form:"score"`
	Directors    string        `json:"directors" form:"directors"`
	Actors       string        `json:"actors" form:"actors"`
	Theaters     []string      `json:"theaters" form:"theaters"`
	TheaterId    uint          `json:"theater_id" form:"theater_id"`
	model.BasePage
}

// Create 上传新电影
func (service *MovieService) Create(ctx context.Context, movieImg, directorImg, actorImg []*multipart.FileHeader) serializer.Response {
	var directors []model.Director
	var actors []model.Actor
	var err error
	code := e.Success

	if len(movieImg) == 0 {
		code = e.ErrorMovieIndex
		util.LogrusObj.Infoln("ErrorMovieIndex", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//以第一张作为封面图
	tmp, _ := movieImg[0].Open()
	path, err := UploadMovieIndexToLocalStatic(tmp, service.ChineseName)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("UploadMovieToLocalStatic", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	directorSlice := strings.Split(service.Directors, ",")
	actorSlice := strings.Split(service.Actors, ",")

	//整理导演演员信息
	for _, director := range directorSlice {
		directors = append(directors, model.Director{Name: director, ImageURL: conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.DirectorPath + director + ".jpg"})
	}

	for _, actor := range actorSlice {
		actors = append(actors, model.Actor{Name: actor, ImageURL: conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.ActorPath + actor + ".jpg"})
	}

	movie := &model.Movie{
		ChineseName:  service.ChineseName,
		EnglishName:  service.EnglishName,
		CategoryId:   service.CategoryIds,
		Area:         service.Area,
		Duration:     service.Duration,
		ShowTime:     service.ShowTime,
		Introduction: service.Introduction,
		ImgPath:      path,
		OnSale:       false,
		Score:        service.Score,
		Directors:    directors,
		Actors:       actors,
	}

	MovieDao := dao.NewMovieDao(ctx)
	err = MovieDao.CreateMovie(movie)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateMovie", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//上传详情图片
	_, err = UploadMovieToLocalStatic(movieImg, service.ChineseName)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//上传导演图片
	_, err = UploadDirectorToLocalStatic(directorImg, directorSlice)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//上传演员图片
	_, err = UploadActorToLocalStatic(actorImg, actorSlice)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildMovie(movie),
	}
}

// ListAll 获取全部电影列表
func (service *MovieService) ListAll(ctx context.Context) serializer.Response {
	var movies []*model.Movie
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	categoryId := uint(0)
	if len(service.CategoryId) != 0 {
		categoryId = service.CategoryId[0]
	}
	productDao := dao.NewMovieDao(ctx)
	total, err := productDao.CountMovieByCondition(categoryId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountMovieByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewMovieDaoByDB(productDao.DB)
		movies, _ = productDao.ListMovieByCondition(categoryId, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(total))
}

// ListHot 获取热映电影列表
func (service *MovieService) ListHot(ctx context.Context) serializer.Response {
	var movies []*model.Movie
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	categoryId := uint(0)
	if len(service.CategoryId) != 0 {
		categoryId = service.CategoryId[0]
	}
	productDao := dao.NewMovieDao(ctx)
	total, err := productDao.CountHotMovieByCondition(categoryId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountMovieByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewMovieDaoByDB(productDao.DB)
		movies, _ = productDao.ListHotMovieByCondition(categoryId, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(total))
}

// ListHotByTheater 获取影院热映电影列表
func (service *MovieService) ListHotByTheater(ctx context.Context) serializer.Response {
	var err error
	var movies []*model.Movie
	code := e.Success

	movieDao := dao.NewMovieDao(ctx)
	total, err := movieDao.CountHotMovieByTheater(service.TheaterId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountHotMovieByTheater", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		movies, _ = movieDao.ListHotMovieByTheater(service.TheaterId)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(total))
}

// ListUnreleased 获取未上映电影列表
func (service *MovieService) ListUnreleased(ctx context.Context) serializer.Response {
	var movies []*model.Movie
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	categoryId := uint(0)
	if len(service.CategoryId) != 0 {
		categoryId = service.CategoryId[0]
	}
	productDao := dao.NewMovieDao(ctx)
	total, err := productDao.CountUnreleasedMovieByCondition(categoryId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountMovieByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewMovieDaoByDB(productDao.DB)
		movies, _ = productDao.ListUnreleasedMovieByCondition(categoryId, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(total))
}

// ListSales 获取电影票房列表
func (service *MovieService) ListSales(ctx context.Context) serializer.Response {
	var movies []*model.Movie
	code := e.Success
	productDao := dao.NewMovieDao(ctx)
	movies, err := productDao.ListMovieBySales(service.BasePage)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("ListMovieBySales", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(len(movies)))
}

// Search 搜索电影
func (service *MovieService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	var movies []*model.Movie

	productDao := dao.NewMovieDao(ctx)
	//精确查找
	movie, err := productDao.SearchMovieExactly(service.Introduction)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	movies = append(movies, movie)
	//模糊
	name := `%`
	for _, ch := range service.Introduction {
		name += string(ch) + "%"
	}
	movies, err = productDao.SearchMovie(name, service.BasePage)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(len(movies)))
}

// ListIndexHotMovies 获取首页热映电影
func (service *MovieService) ListIndexHotMovies(ctx context.Context) serializer.Response {
	var movies []*model.Movie
	var err error
	code := e.Success

	service.PageSize = 8

	now := time.Now()
	// 获取30天前的日期
	preDate := now.AddDate(0, 0, -30)

	productDao := dao.NewMovieDao(ctx)
	total, err := productDao.CountIndexHotMovie(now, preDate)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountIndexHotMovie", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewMovieDaoByDB(productDao.DB)
		movies, _ = productDao.ListIndexHotMovie(now, preDate, service.PageSize)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(total))
}

// Delete 删除电影
func (service *MovieService) Delete(ctx context.Context) serializer.Response {
	code := e.Success

	movieDao := dao.NewMovieDao(ctx)
	movie, err := movieDao.GetMovieByMovieID(service.MovieId)
	if err != nil {
		code = e.ErrorMovieId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if movie.OnSale {
		code = e.ErrorMovieStatus
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	_, err = movieDao.DeleteMovie(service.MovieId)
	if err != nil {
		code = e.ErrorMovieId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Get 获取电影详细信息
func (service *MovieService) Get(ctx context.Context) serializer.Response {
	code := e.Success
	movieDao := dao.NewMovieDao(ctx)
	movie, err := movieDao.GetMovieByMovieID(service.MovieId)
	if err != nil {
		code = e.ErrorMovieId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println(movie)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildMovie(movie),
	}
}
