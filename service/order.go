package service

import (
	"TTMS_Web/serializer"
	"context"
)

type OrderService struct {
	UserID    uint    `json:"user_id" form:"user_id"`
	MovieID   uint    `json:"movie_id" form:"movie_id"`
	AddressID uint    `json:"address_id" form:"address_id"`
	ThreatID  []uint  `json:"threat_id" form:"threat_id"`
	Num       uint    `json:"num" form:"num"`
	ShowTime  string  `json:"show_time" form:"show_time" time_format:"2006-01-02 15:04"`
	Type      uint    `json:"type" form:"type"`
	Money     float64 `json:"money" form:"money"`
}

func (service *OrderService) Submit(ctx context.Context) serializer.Response {
	return serializer.Response{}
}
