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
	"time"
)

type OrderService struct {
	OrderID   uint    `form:"order_id" json:"order_id"`
	MovieID   uint    `json:"movie_id" form:"movie_id"`
	SessionID uint    `json:"session_id" form:"session_id"`
	ThreatID  uint    `json:"threat_id" form:"threat_id"`
	Seat      string  `json:"seat" form:"seat"`
	Num       int     `json:"num" form:"num"`
	Type      uint    `json:"type" form:"type"`
	Money     float64 `json:"money" form:"money"`
	model.BasePage
}

// Submit 提交订单逻辑
func (service *OrderService) Submit(ctx context.Context, userID uint) serializer.Response {
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

	// 从缓存获取场次信息
	session, err := cache.GetSessionInfo(ctx, rdb, service.SessionID)
	if err != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}
	//判断座位是否有人占用
	if util.IsRepeatSeat(service.Seat, session.SeatStatus, session.SeatRow) {
		return serializer.Response{
			Status: e.ErrorSeat,
			Msg:    e.GetMsg(e.ErrorSeat),
		}
	}
	// 更新 session 的座位和余票
	util.UpdateSessionSeat(session, service.Seat, service.Num)
	// 序列化 session
	sessionByte, err := json.Marshal(session)
	if err != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}
	// 更新场次信息
	sessionDao := dao.NewSessionDaoByDB(txDB)
	if err := sessionDao.UpdateSessionByID(service.SessionID, session); err != nil {
		txDB.Rollback()
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	// 创建订单
	order = &model.Order{
		UserID:    userID,
		MovieID:   service.MovieID,
		SessionID: service.SessionID,
		TheaterID: session.TheaterID,
		Seat:      service.Seat,
		Num:       service.Num,
		Type:      0,
		Money:     session.Price,
	}

	orderDao := dao.NewOrderDaoByDB(txDB)
	if order, err = orderDao.AddOrder(order); err != nil {
		txDB.Rollback()
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	// 提交数据库事务
	if err := txDB.Commit().Error; err != nil {
		txDB.Rollback()
		_ = cache.AlterStock(ctx, rdb, service.SessionID, session.SurplusTicket+service.Num)
		_ = cache.DelSessionInfo(ctx, rdb, service.SessionID)
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}

	// 更新 Redis 缓存
	pipe := rdb.TxPipeline()
	cache.AlterStockPipe(ctx, pipe, service.SessionID, session.SurplusTicket)
	cache.SetSessionInfoPipe(ctx, pipe, string(sessionByte), service.SessionID)
	_, err = pipe.Exec(ctx)
	if err != nil {
		txDB.Rollback()
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(e.Error),
		}
	}
	endTime := time.Now().Add(14 * time.Minute)
	endString := endTime.Format("2006-01-02 15:04:05")
	err = cache.SetOrderCount(ctx, rdb, endString, order.ID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(e.Error),
		}
	}
	go startCountdown(order.ID, ctx)
	session = &model.Session{}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order),
	}
}

// 开始倒计时
func startCountdown(orderID uint, ctx context.Context) {
	rdb := cache.GetRedisClient()
	time.Sleep(14 * time.Minute)
	orderDao := dao.NewOrderDao(ctx)
	order, _ := orderDao.GetOrderByOrderID(orderID)
	//未支付订单
	if order.Type == 0 {
		session, _ := cache.GetSessionInfo(ctx, rdb, order.SessionID)
		_ = cache.AlterStock(ctx, rdb, order.SessionID, session.SurplusTicket-order.Num)
		_ = cache.DelSessionInfo(ctx, rdb, order.SessionID)
		_ = orderDao.DeleteOrderByID(orderID)
	}
}

// Confirm 确认订单(查看)订单时间
func (service *OrderService) Confirm(ctx context.Context) serializer.Response {
	code := e.Success
	rdb := cache.GetRedisClient()
	order, err := cache.GetOrderInfo(ctx, rdb, service.OrderID)
	if err != nil {
		code = e.ErrorOrderID
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(e.Error),
		}
	}
	endString, err := cache.GetOrderCount(ctx, rdb, service.OrderID)
	if err != nil {
		code = e.ErrorEndTime
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endString, time.Local)
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
		Data:   serializer.BuildOrderWithTime(order, endTime.Sub(time.Now()).Seconds()),
	}
}

// Pay 支付订单逻辑
func (service *OrderService) Pay(ctx context.Context, userID uint) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order := &model.Order{
		Type: 1,
	}
	err := orderDao.UpdateOrderByID(service.OrderID, order)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(userID)
	if err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if user.Money-order.Money < 0 {
		code = e.ErrorUserMoney
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Money -= order.Money
	err = userDao.UpdateUserByID(userID, user)
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
	}
}

// Get 获取该用户订单
func (service *OrderService) Get(ctx context.Context, userID uint) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	userDao := dao.NewUserDao(ctx)
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	//判断用户是否存在
	_, err := userDao.GetUserByID(userID)
	if err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	orders, err := orderDao.ListUserOrders(userID, service.BasePage)
	if err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println(orders)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrders(orders),
	}
}

// Return 退票逻辑
func (service *OrderService) Return(ctx context.Context, userID uint) serializer.Response {
	session := &model.Session{}
	code := e.Success
	rdb := cache.GetRedisClient()
	orderDao := dao.NewOrderDao(ctx)

	order, err := orderDao.GetOrderByOrderID(service.OrderID)
	if err != nil {
		code = e.ErrorOrderID
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if order.Type == 0 {
		code = e.ErrorOrderType
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = cache.AlterStock(ctx, rdb, service.SessionID, session.SurplusTicket-service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = cache.DelSessionInfo(ctx, rdb, service.SessionID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = orderDao.DeleteOrderByID(service.OrderID)
	if err != nil {
		code = e.ErrorOrderID
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(userID)
	if err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Money += order.Money
	err = userDao.UpdateUserByID(userID, user)
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
	}

}
