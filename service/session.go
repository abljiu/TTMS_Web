package service

import (
	"TTMS_Web/cache"
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"sync"
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

type SessionListRequest struct {
	TheaterID uint   `form:"theater_id" json:"theater_id" binding:"required"`
	MovieID   uint   `form:"movie_id" json:"movie_id"`
	Date      string `json:"date" form:"date"`
	model.BasePage
}

type SessionRequest struct {
	ID uint `json:"id" form:"id"`
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
	if sessionDao.IsTimeOverlap(session.ShowTime, session.EndTime) {
		code = e.ErrorSessionTime
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
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
	}

	theaterDao := dao.NewTheaterDao(ctx)
	theater, err := theaterDao.GetTheaterByID(service.TheaterID)
	if err != nil {
		code = e.ErrorTheaterID
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	movie.Theaters = append(movie.Theaters, *theater)
	err = movieDao.UpdateMovie(movie.ID, movie)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//添加库存
	err = cache.InitializeStock(ctx, rdb, session.ID, hall.SeatNum)
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

// List 获取场次列表
func (service *SessionListRequest) List(ctx context.Context) serializer.Response {
	var sessions []*model.Session
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewSessionDao(ctx)

	//curTime
	curTime := time.Now()
	curTimeStr := curTime.Format("2006-01-02 15:04")

	var total int64
	if len(service.Date) != 0 && service.MovieID != 0 { //date && movieID
		total, err = productDao.CountSessionByMovieIDAndDate(service.TheaterID, service.MovieID, service.Date, curTimeStr)
	} else if len(service.Date) != 0 { //date
		total, err = productDao.CountSessionByDate(service.TheaterID, service.Date, curTimeStr)
	} else if service.MovieID != 0 { //movieID
		total, err = productDao.CountSessionByMovieID(service.TheaterID, service.MovieID, curTimeStr)
	} else {
		total, err = productDao.CountSession(service.TheaterID, curTimeStr)
	}
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountSession", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewSessionDaoByDB(productDao.DB)
		if len(service.Date) != 0 && service.MovieID != 0 { //date && movieID
			sessions, _ = productDao.ListSessionByDateAndMovieID(service.TheaterID, service.MovieID, service.Date, curTimeStr, service.BasePage)
		} else if len(service.Date) != 0 { //date
			sessions, _ = productDao.ListSessionByDate(service.TheaterID, service.Date, curTimeStr, service.BasePage)
		} else if service.MovieID != 0 { //movieID
			sessions, _ = productDao.ListSessionByMovieID(service.TheaterID, service.MovieID, curTimeStr, service.BasePage)
		} else {
			sessions, _ = productDao.ListSession(service.TheaterID, curTimeStr, service.BasePage)
		}
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildSessions(sessions, service.TheaterID, ctx), uint(total))
}

// Get 获取某场次详细信息
func (service *SessionRequest) Get(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	SessionDao := dao.NewSessionDao(ctx)
	session, err := SessionDao.GetSessionByID(service.ID)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("GetSession", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	movieDao := dao.NewMovieDao(ctx)
	hallDao := dao.NewHallDao(ctx)
	//movie
	movie, err := movieDao.GetMovieByMovieID(session.MovieID)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("GetMovieByMovieID", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//hall
	hall, err := hallDao.GetHallByHallID(session.HallID)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("GetHallByHallID", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildSession(session, movie, hall),
	}
}
