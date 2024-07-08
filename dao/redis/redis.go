package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"mxaof_blog/dao/global"
)

var GlobalRedis *redis.Client

func RedisInit() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.GlobalConfig.Redis.Address,
		Password: global.GlobalConfig.Redis.Password,
		DB:       global.GlobalConfig.Redis.Db,
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
