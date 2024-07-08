package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var GlobalRedis *redis.Client

func RedisInit() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.95.199.174:6380",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
	GlobalRedis = rdb

	_, err := GlobalRedis.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	_ = GlobalRedis.Close()
}
