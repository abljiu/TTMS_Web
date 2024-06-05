package serializer

import (
	"TTMS_Web/model"
	"TTMS_Web/pkg/util"
	"time"
)

type Order struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Movie       string    `json:"movie"`
	Theater     string    `json:"theater"`
	SessionID   uint      `json:"session_id"`
	Seat        [][2]int  `json:"seat"`
	Num         int       `json:"num"`
	Type        uint      `json:"type"`
	Money       float64   `json:"money"`
	SurplusTime float64   `json:"surplus_time"`
	ShowTime    time.Time `json:"show_time"`
	Hall        string    `json:"hall"`
}

func BuildOrder(order *model.Order, movie, theater, hall string) *Order {
	seat := make([][2]int, 0)
	seats := util.ParseSeat(order.Seat)
	for i, j := 0, 1; j < len(seats); i, j = i+2, j+2 {
		seat = append(seat, [2]int{seats[i], seats[j]})
	}
	return &Order{
		ID:        order.ID,
		UserID:    order.UserID,
		Movie:     movie,
		Theater:   theater,
		SessionID: order.SessionID,
		Seat:      seat,
		Num:       order.Num,
		Type:      order.Type,
		Money:     order.Money,
		ShowTime:  order.Session.ShowTime,
		Hall:      hall,
	}
}

func BuildOrderWithTime(order *model.Order, surplusTime float64, movie, theater, hall string) *Order {
	seat := make([][2]int, 0)
	seats := util.ParseSeat(order.Seat)
	for i, j := 0, 1; j < len(seats); i, j = i+2, j+2 {
		seat = append(seat, [2]int{seats[i], seats[j]})
	}
	return &Order{
		ID:          order.ID,
		UserID:      order.UserID,
		Movie:       movie,
		Theater:     theater,
		SessionID:   order.SessionID,
		Seat:        seat,
		Num:         order.Num,
		Type:        order.Type,
		Money:       order.Money,
		SurplusTime: surplusTime,
		ShowTime:    order.Session.ShowTime,
		Hall:        hall,
	}
}

func BuildOrders(items []*model.Order, movies, theaters, halls []string) (products []Order) {
	for i := 0; i < len(items); i++ {
		product := BuildOrder(items[i], movies[i], theaters[i], halls[i])
		products = append(products, *product)
	}
	return products
}
