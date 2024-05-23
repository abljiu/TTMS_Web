package service

import (
	"TTMS_Web/conf"
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"fmt"
	"mime/multipart"
	"sync"
)

type MovieService struct {
	MovieId      uint     `json:"movie_id" form:"movie_id"`
	ChineseName  string   `json:"chinese_name" form:"chinese_name"`
	EnglishName  string   `json:"english_name" form:"english_name"`
	CategoryId   []uint   `json:"category_id" form:"category_id"`
	Area         string   `json:"area" form:"area"`
	Duration     string   `json:"duration" form:"duration"`
	ShowTime     string   `json:"show_time" form:"show_time" time_format:"2006-01-02"`
	Introduction string   `json:"introduction" form:"introduction"`
	OnSale       bool     `json:"on_sale" form:"on_sale"`
	Score        float64  `json:"score" form:"score"`
	Directors    []string `json:"directors" form:"directors"`
	Actors       []string `json:"actors" form:"actors"`
	model.BasePage
}

// Create 上传新电影
func (service *MovieService) Create(ctx context.Context, uid uint, movieImg, directorImg, actorImg []*multipart.FileHeader) serializer.Response {
	var directors []model.Director
	var actors []model.Actor
	var err error
	code := e.Success

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

	//整理导演演员信息
	for _, director := range service.Directors {
		directors = append(directors, model.Director{Name: director, ImageURL: conf.Config_.Path.Host + conf.Config_.Path.DirectorPath})
	}
	for _, actor := range service.Actors {
		actors = append(actors, model.Actor{Name: actor, ImageURL: conf.Config_.Path.Host + conf.Config_.Path.ActorPath})
	}

	movie := &model.Movie{
		ChineseName: service.ChineseName,
		EnglishName: service.EnglishName,
		CategoryId:  service.CategoryId,
		Area:        service.Area,
		Duration:    service.Duration,
		ImgPath:     path,
		OnSale:      false,
		Score:       service.Score,
		Directors:   directors,
		Actors:      actors,
	}
	MovieDao := dao.NewMovieDao(ctx)
	err = MovieDao.CreateMovie(movie)
	//bug 多列
	fmt.Println(err)
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
	_, err = UploadDirectorToLocalStatic(directorImg, service.Directors)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//上传演员图片
	_, err = UploadActorToLocalStatic(actorImg, service.Actors)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//根据categoryId返回string
	categoryDao := dao.NewCategoryDao(ctx)
	categoryString, err := categoryDao.GetCategory(service.CategoryId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("GetCategory", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildMovie(movie, categoryString),
	}
}

// List 获取电影列表
func (service *MovieService) List(ctx context.Context) serializer.Response {
	var products []*model.Movie
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryId[0] != 0 {
		condition["category_id"] = service.CategoryId[0]
	}
	productDao := dao.NewMovieDao(ctx)
	total, err := productDao.CountMovieByCondition(condition)
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
		products, _ = productDao.ListMovieByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()
	categoryDao := dao.NewCategoryDao(ctx)
	var categoryStrings []string
	for _, product := range products {
		categoryString, err := categoryDao.GetCategory(product.CategoryId)
		if err != nil {
			code = e.Error
			util.LogrusObj.Infoln("GetCategory", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		categoryStrings = append(categoryStrings, categoryString)
	}

	return serializer.BuildListResponse(serializer.BuildMovies(products, categoryStrings), uint(total))
}

// Search 搜索电影
func (service *MovieService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewMovieDao(ctx)
	movies, err := productDao.SearchMovie(service.Introduction, service.BasePage)
	if err != nil {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	categoryDao := dao.NewCategoryDao(ctx)
	var categoryStrings []string
	for _, movie := range movies {
		categoryString, err := categoryDao.GetCategory(movie.CategoryId)
		if err != nil {
			code = e.Error
			util.LogrusObj.Infoln("GetCategory", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		categoryStrings = append(categoryStrings, categoryString)
	}
	return serializer.BuildListResponse(serializer.BuildMovies(movies, categoryStrings), uint(len(movies)))
}
