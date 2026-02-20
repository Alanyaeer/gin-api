package demo

import (
	"chat-system/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)
var redisCfg = config.Cfg.Redis
func DemoRedisOps() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: redisCfg.Addr,
		DB:   redisCfg.DB,
	})
	slog.Info(fmt.Sprintf("rdb type is %T", rdb))
	rdb.Set(ctx, "key", "Ciallo", time.Hour)
}

func DemoRedisGetVal() string {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: redisCfg.Addr,
		DB:   redisCfg.DB,
	})
	result, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		slog.Error("Failed to get value from Redis", "error", err)
		return ""
	}
	return result
}

func DemoTimeOutRedisOps() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
		DB:   config.RedisDB,
	})
	err := rdb.Set(ctx, "key", "Ciallo", time.Hour).Err()
	if err != nil {
		slog.Error("Failed to set value in Redis", "error", err)
	} else {
		slog.Info("Successfully set value in Redis")
	}
}
