package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type TheaterService struct {
	TheaterId   uint   `json:"theater_id" form:"theater_id"`
	TheaterName string `json:"name" form:"name" `
	Address     string `json:"address" form:"address"`
	CreateTime  string `json:"create_time" form:"create_time" time_format:"2006-01-02"`
	HallNum     int    `json:"hall_num" form:"hall_num"`
	model.BasePage
}

// Create 上传新剧院
func (service *TheaterService) Create(ctx context.Context) serializer.Response {
	var err error
	code := e.Success

	Theater := &model.Theater{
		Model:   gorm.Model{},
		Name:    service.TheaterName,
		Address: service.Address,
		HallNum: service.HallNum,
	}

	TheaterDao := dao.NewTheaterDao(ctx)
	err = TheaterDao.CreateTheater(Theater)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateTheater", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTheater(Theater),
	}
}

// List 获取剧院列表
func (service *TheaterService) List(ctx context.Context) serializer.Response {
	var Theaters []*model.Theater
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewTheaterDao(ctx)
	total, err := productDao.CountTheater()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CountTheaterByCondition", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewTheaterDaoByDB(productDao.DB)
		Theaters, _ = productDao.ListTheater(service.BasePage)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("success")
	return serializer.BuildListResponse(serializer.BuildTheaters(Theaters), uint(total))
}

// Search 搜索剧院根据名称
func (service *TheaterService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewTheaterDao(ctx)
	Theaters, err := productDao.SearchTheater(service.TheaterName, service.BasePage)
	if err != nil {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildTheaters(Theaters), uint(len(Theaters)))
}

// Search 搜索剧院根据名称
func (service *TheaterService) SearchById(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewTheaterDao(ctx)
	Theater, err := productDao.GetTheaterByID(service.TheaterId)
	if err != nil {
		util.LogrusObj.Infoln("SearchProduct", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTheater(Theater),
	}
}

// Update 修改信息
func (service *TheaterService) Update(ctx context.Context, uid uint) serializer.Response {
	var err error
	code := e.Success
	TheaterDao := dao.NewTheaterDao(ctx)
	Theater := &model.Theater{
		Model:   gorm.Model{},
		Name:    service.TheaterName,
		Address: service.Address,
		HallNum: service.HallNum,
	}
	err = TheaterDao.UpdateTheaterByID(uid, Theater)
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
		Data:   serializer.BuildTheater(Theater),
	}
}

// Delete 删除信息
func (service *TheaterService) Delete(ctx context.Context, uid uint) serializer.Response {
	var Theater *model.Theater
	var err error
	code := e.Success
	TheaterDao := dao.NewTheaterDao(ctx)
	Theater, err = TheaterDao.GetTheaterByID(uid)
	_, err = TheaterDao.DeleteTheaterByID(uid)
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
		Data:   serializer.BuildTheater(Theater),
	}
}
