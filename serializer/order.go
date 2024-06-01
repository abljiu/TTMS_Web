package serializer

import (
	"TTMS_Web/model"
	"TTMS_Web/pkg/util"
)

type Order struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	MovieID   uint      `json:"movie_id"`
	TheaterID uint      `json:"theater_id"`
	SessionID uint      `json:"session_id"`
	Seat      [][2]uint `json:"seat"`
	Num       int       `json:"num"`
	Type      uint      `json:"type"`
	Money     float64   `json:"money"`
}

func BuildOrder(order *model.Order) *Order {
	seat := make([][2]uint, 0)
	seats := util.ParseSeat(order.Seat)
	for i, j := 0, 1; j < len(seats); i, j = i+1, j+1 {
		seat = append(seat, [2]uint{seats[i], seats[j]})
	}
	return &Order{
		ID:        order.ID,
		UserID:    order.UserID,
		MovieID:   order.MovieID,
		TheaterID: order.TheaterID,
		SessionID: order.SessionID,
		Seat:      seat,
		Num:       order.Num,
		Type:      order.Type,
		Money:     order.Money,
	}
}