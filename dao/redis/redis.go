package redis

import (
	"context"
	"fmt"
	"gin-api/settings"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(ctx).Result()
	return
}
func Close() {
	err := rdb.Close()
	if err != nil {
		return
	}
}
