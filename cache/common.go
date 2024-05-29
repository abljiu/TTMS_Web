package cache

import (
	"TTMS_Web/conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var RedisClient *redis.Client

func InitRedis() {
	Redis()
}

func Redis() {
	db, _ := strconv.ParseUint(conf.Config_.Redis.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Config_.Redis.RedisAddr,
		DB:       int(db),
		Password: conf.Config_.Redis.RedisPw,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	RedisClient = client
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
