package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// InitializeStock 初始化场次库存
func InitializeStock(ctx context.Context, rdb *redis.Client, sessionID uint, stock uint) error {
	key := fmt.Sprintf("ticket_stock:%d", sessionID)
	err := rdb.Set(ctx, key, stock, 0).Err()
	return err
}

// AlterStock 更改场次库存
func AlterStock(ctx context.Context, rdb *redis.Client, sessionID uint, num int) error {
	key := fmt.Sprintf("ticket_stock:%d", sessionID)
	return rdb.Set(ctx, key, num, 0).Err()
}

// AlterStockPipe 更改场次库存
func AlterStockPipe(ctx context.Context, pipe redis.Pipeliner, sessionID uint, num int) {
	key := fmt.Sprintf("ticket_stock:%d", sessionID)
	pipe.Set(ctx, key, num, 0)
}
