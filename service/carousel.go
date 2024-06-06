package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/serializer"
	"context"
	"fmt"
)

type CarouselService struct {
	MovieID uint `form:"movie_id" json:"movie_id" `
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println("asdasdasd", carousels)
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}

// Add 添加
func (service *CarouselService) Add(ctx context.Context) serializer.Response {
	code := e.Success

	movieDao := dao.NewMovieDao(ctx)
	movie, err := movieDao.GetMovieByMovieID(service.MovieID)
	if err != nil {
		code = e.ErrorMovieId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	carousel := &model.Carousel{
		MovieID: movie.ID,
		Movie:   *movie,
	}
	carouselDao := dao.NewCarouselDao(ctx)
	err = carouselDao.CreateCarousel(carousel)
	if err != nil {
		code = e.Error
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
