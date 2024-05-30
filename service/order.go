package service

import (
	"TTMS_Web/cache"
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"encoding/json"
)

type OrderService struct {
	UserID    uint    `json:"user_id" form:"user_id"`
	MovieID   uint    `json:"movie_id" form:"movie_id"`
	SessionID uint    `json:"session_id" form:"session_id"`
	ThreatID  uint    `json:"threat_id" form:"threat_id"`
	Seat      string  `json:"seat" form:"seat"`
	Num       int     `json:"num" form:"num"`
	Type      uint    `json:"type" form:"type"`
	Money     float64 `json:"money" form:"money"`
}

// Submit 提交订单逻辑
func (service *OrderService) Submit(ctx context.Context) serializer.Response {
	session := &model.Session{}
	code := e.Success
	rdb := cache.GetRedisClient()
	sessionInfo, err := cache.GetSessionInfo(ctx, rdb, service.SessionID)
	//反序列化到结构体
	err = json.Unmarshal([]byte(sessionInfo), session)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//根据string 修改session seat 和 surplusTicket
	util.UpdateSessionSeat(session, service.Seat, service.Num)
	//序列化
	sessionByte, err := json.Marshal(session)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	sessionInfo = string(sessionByte)
	db := dao.NewDBClient(ctx)
	tx := db.Begin()
	if tx.Error != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	pipe := rdb.TxPipeline()
	cache.AlterStockPipe(ctx, pipe, service.SessionID, session.SurplusTicket)
	cache.SetSessionInfoPipe(ctx, pipe, sessionInfo, service.SessionID)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return serializer.Response{
			Status: e.ErrorCacheAddSession,
			Msg:    e.GetMsg(code),
		}
	}

	//更新场次信息
	sessionDao := dao.NewSessionDaoByDB(tx)
	err = sessionDao.UpdateSessionByID(service.SessionID, session)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建订单
	order := &model.Order{
		UserID:    service.UserID,
		MovieID:   service.MovieID,
		SessionID: service.SessionID,
		TheaterID: session.TheaterID,
		Seat:      service.Seat,
		Num:       service.Num,
		Type:      0,
		Money:     session.Price,
	}
	orderDao := dao.NewOrderDaoByDB(tx)
	err = orderDao.AddOrder(order)
	if err != nil {
		code = e.ErrorAddOrder
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback() // 如果提交失败，回滚事务
		code = e.ErrorAddOrder
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order),
	}
}

func (service *OrderService) Return(ctx context.Context) serializer.Response {
	return serializer.Response{}
}
