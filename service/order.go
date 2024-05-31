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
	"time"
)

type OrderService struct {
	OrderID   uint    `form:"order_id" json:"order_id"`
	UserID    uint    `json:"user_id" form:"user_id"`
	MovieID   uint    `json:"movie_id" form:"movie_id"`
	SessionID uint    `json:"session_id" form:"session_id"`
	ThreatID  uint    `json:"threat_id" form:"threat_id"`
	Seat      string  `json:"seat" form:"seat"`
	Num       int     `json:"num" form:"num"`
	Type      uint    `json:"type" form:"type"`
	Money     float64 `json:"money" form:"money"`
}

var sessionWithMutex model.SessionWithMutex

// Submit 提交订单逻辑 //todo 逻辑问题
func (service *OrderService) Submit(ctx context.Context) serializer.Response {
	var err error
	var code = e.Success
	order := &model.Order{}
	rdb := cache.GetRedisClient()
	// 创建数据库事务
	db := dao.NewDBClient(ctx)
	txDB := db.Begin()
	if txDB.Error != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	//获取读写锁
	sessionWithMutex.Mutex.Lock()
	defer sessionWithMutex.Mutex.Unlock()

	// 从缓存获取场次信息
	sessionWithMutex.Session, err = cache.GetSessionInfo(ctx, rdb, service.SessionID)
	if err != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}
	//判断座位是否有人占用
	if util.IsRepeatSeat(service.Seat, sessionWithMutex.Session.SeatStatus, sessionWithMutex.Session.SeatRow) {
		return serializer.Response{
			Status: e.ErrorSeat,
			Msg:    e.GetMsg(e.ErrorSeat),
		}
	}
	// 更新 session 的座位和余票
	util.UpdateSessionSeat(sessionWithMutex.Session, service.Seat, service.Num)
	// 序列化 session
	sessionByte, err := json.Marshal(sessionWithMutex.Session)
	if err != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}
	// 更新场次信息
	sessionDao := dao.NewSessionDaoByDB(txDB)
	if err := sessionDao.UpdateSessionByID(service.SessionID, sessionWithMutex.Session); err != nil {
		txDB.Rollback()
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	// 创建订单
	order = &model.Order{
		UserID:    service.UserID,
		MovieID:   service.MovieID,
		SessionID: service.SessionID,
		TheaterID: sessionWithMutex.Session.TheaterID,
		Seat:      service.Seat,
		Num:       service.Num,
		Type:      0,
		Money:     sessionWithMutex.Session.Price,
	}

	orderDao := dao.NewOrderDaoByDB(txDB)
	if err := orderDao.AddOrder(order); err != nil {
		txDB.Rollback()
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	// 提交数据库事务
	if err := txDB.Commit().Error; err != nil {
		txDB.Rollback()
		_ = cache.AlterStock(ctx, rdb, service.SessionID, sessionWithMutex.Session.SurplusTicket+service.Num)
		_ = cache.DelSessionInfo(ctx, rdb, service.SessionID)
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	// 更新 Redis 缓存
	pipe := rdb.TxPipeline()
	cache.AlterStockPipe(ctx, pipe, service.SessionID, sessionWithMutex.Session.SurplusTicket)
	cache.SetSessionInfoPipe(ctx, pipe, string(sessionByte), service.SessionID)
	_, err = pipe.Exec(ctx)
	if err != nil {
		txDB.Rollback()
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}
	sessionWithMutex.Session = &model.Session{}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order),
	}
}

// 写入倒计时
func startCountdown(orderID uint, ctx context.Context) {
	endTime := time.Now().Add(14 * time.Minute)
	rdb := cache.GetRedisClient()
	endString := endTime.Format("2006-01-02 15:04:05")
	cache.SetOrderCount(ctx, rdb, endString, orderID)
}

// Confirm 确认订单(查看)订单时间
func (service *OrderService) Confirm(ctx context.Context) serializer.Response {
	code := e.Success
	rdb := cache.GetRedisClient()
	endString, err := cache.GetOrderCount(ctx, rdb, service.OrderID)
	if err != nil {
		code = e.ErrorEndTime
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", endString)
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
		Data:   endTime.Sub(time.Now()),
	}
}

// Pay 支付订单逻辑
func (service *OrderService) Pay(ctx context.Context) serializer.Response {
	//code := e.Success
	return serializer.Response{}
}

// Return 退票逻辑
func (service *OrderService) Return(ctx context.Context) serializer.Response {
	//session := &model.Session{}
	//code := e.Success
	//
	//rdb := cache.GetRedisClient()
	//sessionInfo, err := cache.GetSessionInfo(ctx, rdb, service.SessionID)
	////反序列化到结构体
	//err = json.Unmarshal([]byte(sessionInfo), session)
	//
	//if err != nil {
	//	code = e.Error
	//	return serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//	}
	//}
	return serializer.Response{}
}

func (service *OrderService) Get(ctx context.Context) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	//判断订单是否存在
	order, err := orderDao.GetOrderByID(service.OrderID)
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
		Data:   serializer.BuildOrder(order),
	}
}

//查看座位是否有重复 有重复且订单未支付 修改前一个订单的seat 已支付返回失败 todo
//if util.IsRepeatSeat(service.Seat, session.SeatStatus, session.SeatRow) {
//	//从缓存获取订单信息
//	orderDao := dao.NewOrderDao(ctx)
//	order, err := cache.GetOrderInfo(ctx, rdb, service.OrderID)
//	if err != nil {
//		code = e.Error
//		return serializer.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//		}
//	}
//}
