package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// InitializeStock 初始化场次库存
func InitializeStock(ctx context.Context, rdb *redis.Client, sessionID uint, stock int) error {
	key := fmt.Sprintf("ticket_stock:%d", sessionID)
	return rdb.Set(ctx, key, stock, 0).Err()
}
