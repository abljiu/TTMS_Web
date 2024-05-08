package cache

import (
	"TTMS_Web/conf"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func init() {
	Redis()
}

func Redis() {
	db, _ := strconv.ParseUint(conf.Config_.Redis.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: conf.Config_.Redis.RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
