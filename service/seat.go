package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"strconv"
)

type ModifySeatService struct {
	HallID int  `json:"hall_id" form:"hall_id"`
	Row    int  `json:"row" form:"row"`
	Column int  `json:"column" form:"column"`
	Option int8 `json:"option" form:"option"` //1：变为已售  2：变为未售出可选
}

// Modify 修改座位状态
func (service *ModifySeatService) Modify(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	hall := &model.Hall{
		ID: service.HallID,
	}
	HallDao := dao.NewHallDao(ctx)
	err = HallDao.GetHall(hall)

	index := (service.Row-1)*hall.SeatColumn*2 + (service.Column-1)*2
	hall.Seat = hall.Seat[:index] + strconv.Itoa(int(service.Option)) + hall.Seat[index+1:]
	seat := &model.ModifySeat{
		HallID: service.HallID,
		Seat:   hall.Seat,
	}
	seatDao := dao.NewSeatDao(ctx)
	err = seatDao.ModifySeat(seat)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("ModifySeat", err)
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
