package serializer

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"context"
	"time"
)

type Session struct {
	ID            uint
	Movie         Movie
	TheaterID     uint
	Theater       model.Theater
	Hall          *model.Hall
	ShowTime      time.Time
	EndTime       time.Time
	SurplusTicket int
	SeatStatus    string
	Price         float64
	SeatRow       uint
}

func BuildSession(item *model.Session, movie *model.Movie, hall *model.Hall) Session {
	return Session{
		ID:            item.ID,
		Movie:         BuildMovie(movie),
		Hall:          hall,
		ShowTime:      item.ShowTime,
		EndTime:       item.ShowTime.Add(movie.Duration),
		SurplusTicket: item.SurplusTicket,
		SeatStatus:    item.SeatStatus,
		Price:         item.Price,
		SeatRow:       uint(hall.SeatRow),
		Theater:       item.Theater,
		TheaterID:     item.TheaterID,
	}
}

func BuildSessions(items []*model.Session, theaterID uint, ctx context.Context) (products []Session) {
	for i := 0; i < len(items); i++ {
		movieDao := dao.NewMovieDao(ctx)
		hallDao := dao.NewHallDao(ctx)
		//根据id获取电影
		movie, err := movieDao.GetMovieByMovieID(items[i].MovieID)
		if err != nil {
			return nil
		}
		// 根据id获取影厅
		hall, err := hallDao.GetHallByHallID(items[i].HallID)
		if err != nil {
			return nil
		}
		product := BuildSession(items[i], movie, hall)
		product.TheaterID = theaterID
		products = append(products, product)
	}
	return products
}
