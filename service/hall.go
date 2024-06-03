package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type HallListRequest struct {
	TheaterId uint `json:"theater_id" form:"theater_id" binding:"required"`
	model.BasePage
}

type HallCreateRequest struct {
	ID         uint   `json:"id" form:"id"`
	TheaterId  uint   `json:"theater_id" form:"theater_id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	SeatRow    int    `json:"seat_row" form:"seat_row" binding:"required,gt=0"`
	SeatColumn int    `json:"seat_column" form:"seat_column" binding:"required,gt=0"`
	Seat       string `json:"seat" form:"seat" binding:"required"`
}

type HallIDRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type HallUpdateRequest struct {
	ID         uint   `json:"id" form:"id" binding:"required"`
	TheaterId  uint   `json:"theater_id" form:"theater_id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	SeatRow    int    `json:"seat_row" form:"seat_row" binding:"required,gt=0"`
	SeatColumn int    `json:"seat_column" form:"seat_column" binding:"required,gt=0"`
	Seat       string `json:"seat" form:"seat" binding:"required"`
}

// List 获取影厅列表
func (service *HallListRequest) List(ctx context.Context) serializer.Response {
	var halls []*model.Hall
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewHallDao(ctx)
	total, err := productDao.CountHallByTheaterID(service.TheaterId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountHallByTheaterID", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewHallDaoByDB(productDao.DB)
		halls, _ = productDao.ListHallByTheaterID(service.TheaterId, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildHalls(halls), uint(total))
}

// Create 新建影厅
func (service *HallCreateRequest) Create(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	var seatNum int
	var num int
	for k := range service.Seat {
		n, _ := strconv.Atoi(string(service.Seat[k]))
		if n != 0 && n != 1 {
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(e.ErrorInvalidSeatParam),
			}
		}
		if n == 1 {
			seatNum++
		}
		num = k
	}
	if num != service.SeatRow*service.SeatColumn {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(e.ErrorInvalidSeatParam),
		}
	}
	if seatNum > service.SeatRow*service.SeatColumn {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(e.ErrorInvalidSeatParam),
		}
	}

	hall := &model.Hall{
		Name:       service.Name,
		TheaterID:  service.TheaterId,
		SeatRow:    service.SeatRow,
		SeatColumn: service.SeatColumn,
		Seat:       service.Seat,
		SeatNum:    seatNum,
	}
	HallDao := dao.NewHallDao(ctx)
	err = HallDao.CreateHall(hall)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateHall", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   err,
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildHall(hall),
	}
}

// Delete 根据影厅id，删除某影厅
func (service *HallIDRequest) Delete(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	hall := &model.Hall{}
	HallDao := dao.NewHallDao(ctx)
	err = HallDao.DeleteHall(hall)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("DeleteHall", err)
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

// Update 更新影厅信息
func (service *HallUpdateRequest) Update(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	var seatNum int
	var num int
	for k := range service.Seat {
		n, _ := strconv.Atoi(string(service.Seat[k]))
		if n != 0 && n != 1 {
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(e.ErrorInvalidSeatParam),
			}
		}
		if n == 1 {
			seatNum++
		}
		num = k
	}
	if num != service.SeatRow*service.SeatColumn {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(e.ErrorInvalidSeatParam),
		}
	}
	if seatNum > service.SeatRow*service.SeatColumn {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(e.ErrorInvalidSeatParam),
		}
	}

	gormModel := gorm.Model{ID: service.ID}
	hall := &model.Hall{
		Model:      gormModel,
		Name:       service.Name,
		TheaterID:  service.TheaterId,
		SeatRow:    service.SeatRow,
		SeatColumn: service.SeatColumn,
		Seat:       service.Seat,
		SeatNum:    seatNum,
	}
	HallDao := dao.NewHallDao(ctx)
	err = HallDao.UpdateHall(hall)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UpdateHall", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   err,
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildHall(hall),
	}
}

// Get 获取某影厅详细信息
func (service *HallIDRequest) Get(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	gormModel := gorm.Model{ID: service.ID}
	hall := &model.Hall{
		Model: gormModel,
	}
	HallDao := dao.NewHallDao(ctx)
	err = HallDao.GetHall(hall)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("GetHall", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildHall(hall),
	}
}
