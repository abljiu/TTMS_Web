package service

import (
	"TTMS_Web/cache"
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/serializer"
	"context"
	"time"
)

type SessionServer struct {
	SessionID uint      `form:"session_id" json:"session_id" `
	MovieID   uint      `form:"movie_id" json:"movie_id"`
	Price     float64   `form:"price" json:"price"`
	HallID    uint      `form:"hall_id" json:"hall_id"`
	TheaterID uint      `form:"theater_id" json:"theater_id"`
	ShowTime  time.Time `form:"show_time" json:"show_time" time_format:"2006-01-02 15:04"`
}

// Add 添加场次
func (service *SessionServer) Add(ctx context.Context) serializer.Response {
	code := e.Success

	sessionDao := dao.NewSessionDao(ctx)
	movieDao := dao.NewMovieDao(ctx)
	hallDao := dao.NewHallDao(ctx)
	rdb := cache.GetRedisClient()

	//根据id获取电影
	movie, err := movieDao.GetMovieByMovieID(service.MovieID)
	if err != nil {
		code = e.ErrorMovieId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 根据id获取影厅
	hall, err := hallDao.GetHallByHallID(service.HallID)
	if err != nil {
		code = e.ErrorHallId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	session := &model.Session{
		MovieID:       service.MovieID,
		TheaterID:     service.TheaterID,
		HallID:        service.HallID,
		ShowTime:      service.ShowTime,
		EndTime:       service.ShowTime.Add(movie.Duration),
		SurplusTicket: hall.SeatNum,
		SeatStatus:    hall.Seat,
		SeatRow:       hall.SeatRow,
		Price:         service.Price,
	}

	//添加场次
	err = sessionDao.AddSession(session)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//更新电影信息
	if !movie.OnSale {
		movie.OnSale = true
		err = movieDao.UpdateMovie(movie.ID, movie)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	//添加库存
	err = cache.InitializeStock(ctx, rdb, session.ID, uint(hall.SeatNum))
	if err != nil {
		code = e.ErrorInitializeStock
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

// Alter 修改场次
func (service *SessionServer) Alter(ctx context.Context) serializer.Response {
	code := e.Success
	sessionDao := dao.NewSessionDao(ctx)
	movieDao := dao.NewMovieDao(ctx)

	//判断场次是否存在
	session, err := sessionDao.GetSessionByID(service.SessionID)
	if err != nil {
		code = e.ErrorSessionId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	movie, err := movieDao.GetMovieByMovieID(session.MovieID)
	if err != nil {
		code = e.ErrorMovieId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if session.MovieID != 0 {
		session.MovieID = service.MovieID
	}
	if session.TheaterID != 0 {
		session.TheaterID = service.TheaterID
	}
	if session.TheaterID != 0 {
		session.TheaterID = service.HallID
	}
	if !session.ShowTime.IsZero() {
		session.ShowTime = service.ShowTime
		session.EndTime = service.ShowTime.Add(movie.Duration)
	}

	err = sessionDao.UpdateSessionByID(service.SessionID, session)
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

// Delete 删除场次
func (service *SessionServer) Delete(ctx context.Context) serializer.Response {
	code := e.Success
	sessionDao := dao.NewSessionDao(ctx)
	//判断场次是否存在
	_, err := sessionDao.GetSessionByID(service.SessionID)
	if err != nil {
		code = e.ErrorSessionId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	err = sessionDao.DeleteSessionByID(service.SessionID)
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

// Get  获取场次信息
func (service *SessionServer) Get(ctx context.Context) serializer.Response {
	code := e.Success
	sessionDao := dao.NewSessionDao(ctx)
	//判断场次是否存在
	session, err := sessionDao.GetSessionByID(service.SessionID)
	if err != nil {
		code = e.ErrorSessionId
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildSession(session),
	}
}
