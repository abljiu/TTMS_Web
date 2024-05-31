package cache

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// GetSessionInfo 从缓存中获取场次信息
func GetSessionInfo(ctx context.Context, rdb *redis.Client, sessionID uint) (*model.Session, error) {
	session := &model.Session{}
	key := fmt.Sprintf("session_info:%d", sessionID)
	sessionDao := dao.NewSessionDao(ctx)
	sessionInfo, err := rdb.Get(ctx, key).Result()
	//缓存未命中
	if errors.Is(err, redis.Nil) {
		session, err = sessionDao.GetSessionByID(sessionID)
		if err != nil {
			return nil, err
		}
		//序列化session
		sessionInfoJSON, err := json.Marshal(session)
		sessionInfo = string(sessionInfoJSON)
		if err != nil {
			return nil, err
		}
		//写入缓存
		err = SetSessionInfo(ctx, rdb, sessionInfo, sessionID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal([]byte(sessionInfo), session)
	}
	return session, nil
}

// SetSessionInfo 添加场次信息
func SetSessionInfo(ctx context.Context, rdb *redis.Client, sessionInfoJSON string, sessionID uint) (err error) {
	// 将场次信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("session_info:%d", sessionID)
	err = rdb.Set(ctx, key, sessionInfoJSON, 10*time.Minute).Err()
	return
}

// DelSessionInfo 删除场次信息
func DelSessionInfo(ctx context.Context, rdb *redis.Client, sessionID uint) (err error) {
	// 将场次信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("session_info:%d", sessionID)
	err = rdb.Del(ctx, key).Err()
	return
}

// SetSessionInfoPipe 添加场次信息
func SetSessionInfoPipe(ctx context.Context, pipe redis.Pipeliner, sessionInfoJSON string, sessionID uint) {
	// 将场次信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("session_info:%d", sessionID)
	pipe.Set(ctx, key, sessionInfoJSON, 10*time.Minute)
}

// GetOrderInfo 从缓存中获取订单信息
func GetOrderInfo(ctx context.Context, rdb *redis.Client, orderID uint) (order *model.Order, err error) {
	key := fmt.Sprintf("order_info:%d", orderID)
	orderDao := dao.NewOrderDao(ctx)
	orderInfo, err := rdb.Get(ctx, key).Result()
	//缓存未命中
	if errors.Is(err, redis.Nil) {
		order, err := orderDao.GetOrderByID(orderID)
		if err != nil {
			return nil, err
		}
		//序列化order
		orderInfoJSON, err := json.Marshal(order)
		orderInfo = string(orderInfoJSON)
		if err != nil {
			return nil, err
		}
		//写入缓存
		err = SetOrderInfo(ctx, rdb, orderInfo, orderID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(orderInfo), order)
	return order, nil
}

// SetOrderInfo 添加订单信息
func SetOrderInfo(ctx context.Context, rdb *redis.Client, orderInfoJSON string, orderID uint) (err error) {
	// 将信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("order_info:%d", orderID)
	err = rdb.Set(ctx, key, orderInfoJSON, 10*time.Minute).Err()
	return
}

// SetOrderInfoPipe 添加场次信息
func SetOrderInfoPipe(ctx context.Context, pipe redis.Pipeliner, orderInfoJSON string, orderID uint) {
	// 将场次信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("order_info:%d", orderID)
	pipe.Set(ctx, key, orderInfoJSON, 10*time.Minute)
}
