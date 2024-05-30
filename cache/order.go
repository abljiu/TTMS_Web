package cache

import (
	"TTMS_Web/dao"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// GetSessionInfo 从缓存中获取场信息
func GetSessionInfo(ctx context.Context, rdb *redis.Client, sessionID uint) (string, error) {
	key := fmt.Sprintf("session_info:%d", sessionID)
	sessionDao := dao.NewSessionDao(ctx)
	sessionInfo, err := rdb.Get(ctx, key).Result()
	//缓存未命中
	if errors.Is(err, redis.Nil) {
		session, err := sessionDao.GetSessionByID(sessionID)
		if err != nil {
			return "", err
		}
		//序列化session
		sessionInfoJSON, err := json.Marshal(session)
		sessionInfo = string(sessionInfoJSON)
		if err != nil {
			return "", err
		}
		//写入缓存
		err = SetSessionInfo(ctx, rdb, sessionInfo, sessionID)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return sessionInfo, nil
}

// SetSessionInfo 添加场次信息
func SetSessionInfo(ctx context.Context, rdb *redis.Client, sessionInfoJSON string, sessionID uint) (err error) {
	// 将场次信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("session_info:%d", sessionID)
	err = rdb.Set(ctx, key, sessionInfoJSON, 10*time.Minute).Err()
	return
}

// SetSessionInfoPipe 添加场次信息
func SetSessionInfoPipe(ctx context.Context, pipe redis.Pipeliner, sessionInfoJSON string, sessionID uint) {
	// 将场次信息写入Redis缓存，并设置过期时间
	key := fmt.Sprintf("session_info:%d", sessionID)
	pipe.Set(ctx, key, sessionInfoJSON, 10*time.Minute)
}
