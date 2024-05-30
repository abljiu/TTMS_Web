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
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
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

var countdowns = sync.Map{}

// Submit 提交订单逻辑
func (service *OrderService) Submit(ctx context.Context) serializer.Response {
	var (
		err     error
		code    = e.Success
		session *model.Session
		order   *model.Order
	)
	rdb := cache.GetRedisClient()
	key := fmt.Sprintf("session_info:%d", service.SessionID)

	// 开启 Redis 事务
	txn := rdb.Watch(ctx, func(tx *redis.Tx) error {
		// 从缓存获取场次信息
		session, err = cache.GetSessionInfo(ctx, rdb, service.SessionID)
		if err != nil {
			return err
		}

		// 更新 session 的座位和余票
		util.UpdateSessionSeat(session, service.Seat, service.Num)

		// 序列化 session
		sessionByte, err := json.Marshal(session)
		if err != nil {
			return err
		}

		// 创建数据库事务
		db := dao.NewDBClient(ctx)
		txDB := db.Begin()
		if txDB.Error != nil {
			return txDB.Error
		}

		// 确保在发生错误时正确回滚
		defer func() {
			if r := recover(); r != nil {
				txDB.Rollback()
				_ = cache.AlterStock(ctx, rdb, service.SessionID, session.SurplusTicket+service.Num)
				_ = cache.DelSessionInfo(ctx, rdb, service.SessionID)
				code = e.Error
			}
		}()

		// 更新场次信息
		sessionDao := dao.NewSessionDaoByDB(txDB)
		if err := sessionDao.UpdateSessionByID(service.SessionID, session); err != nil {
			txDB.Rollback()
			return err
		}

		// 创建订单
		order = &model.Order{
			UserID:    service.UserID,
			MovieID:   service.MovieID,
			SessionID: service.SessionID,
			TheaterID: session.TheaterID,
			Seat:      service.Seat,
			Num:       service.Num,
			Type:      0,
			Money:     session.Price,
		}

		orderDao := dao.NewOrderDaoByDB(txDB)
		if err := orderDao.AddOrder(order); err != nil {
			txDB.Rollback()
			return err
		}

		// 提交数据库事务
		if err := txDB.Commit().Error; err != nil {
			txDB.Rollback()
			_ = cache.AlterStock(ctx, rdb, service.SessionID, session.SurplusTicket+service.Num)
			_ = cache.DelSessionInfo(ctx, rdb, service.SessionID)
			return err
		}

		// 更新 Redis 缓存
		pipe := tx.TxPipeline()
		cache.AlterStockPipe(ctx, pipe, service.SessionID, session.SurplusTicket)
		cache.SetSessionInfoPipe(ctx, pipe, string(sessionByte), service.SessionID)
		_, err = pipe.Exec(ctx)
		if err != nil {
			txDB.Rollback()
			return err
		}

		return nil
	}, key)
	if txn != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order),
	}
}

// 倒计时函数
func startCountdown(orderID uint, orderDao *dao.OrderDao) {
	endTime := time.Now().Add(14 * time.Minute)
	countdowns.Store(orderID, endTime)
	//rdb := cache.GetRedisClient()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			remaining := time.Until(endTime)
			if remaining <= 0 {
				// 倒计时结束，检查订单状态并更新
				status, _ := orderDao.CheckOrderTypeByID(orderID)
				if status == 0 {
					//更新缓存和数据库的订单状态

				}
				countdowns.Delete(orderID)
				return
			}
			// 更新倒计时
			countdowns.Store(orderID, endTime)
		}
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
