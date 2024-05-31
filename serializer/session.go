package serializer

import (
	"TTMS_Web/model"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	ID          uint          `json:"id"`
	MovieName   string        `json:"movie_name"`
	TheaterName string        `json:"theater_name"`
	HallName    string        `json:"hall_name"`
	Address     string        `json:"address"`
	Telephone   string        `json:"telephone"`
	Duration    time.Duration `json:"duration"`
	ShowTime    time.Time     `json:"show_time"`
	EndTime     time.Time     `json:"end_time"`
	CategoryID  []uint        `json:"category_id"`
	SeatStatus  string        `json:"seat_status"`
	Price       float64       `json:"price"`
	SeatRow     int           `json:"seat_row"`
	SeatColumn  int           `json:"seat_column"`
}

func BuildSession(session *model.Session) *Session {
	CategoryID := make([]uint, len(session.Movie.CategoryId))
	fmt.Println(session.Movie)
	strSlice := strings.Split(session.Movie.CategoryId, ",")
	for i, str := range strSlice {
		num, _ := strconv.ParseUint(str, 10, 64)
		CategoryID[i] = uint(num)
	}
	return &Session{
		ID:          session.ID,
		MovieName:   session.Movie.ChineseName,
		TheaterName: session.Theater.Name,
		HallName:    session.Hall.Name,
		Address:     session.Theater.Address,
		Telephone:   session.Theater.Telephone,
		Duration:    session.Movie.Duration,
		ShowTime:    session.ShowTime,
		EndTime:     session.EndTime,
		CategoryID:  CategoryID,
		SeatStatus:  session.SeatStatus,
		Price:       session.Price,
		SeatRow:     session.SeatRow,
		SeatColumn:  session.Hall.SeatColumn,
	}
}
