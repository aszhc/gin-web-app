package redis

import (
	"fmt"
	"gin-web-app/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PollSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
