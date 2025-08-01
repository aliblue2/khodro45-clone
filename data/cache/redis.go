package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/aliblue2/khodro45/configs"
	"github.com/aliblue2/khodro45/pkg/logging"
	"github.com/redis/go-redis/v9"
)

var redisClient redis.Client

func InitiateRedis(cfg *configs.Config) error {
	var logger = logging.NewLogger(cfg)

	redisClient = *redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout * time.Second,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		MinIdleConns: 500,
	})

	_, err := redisClient.Ping(context.Background()).Result()

	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
		return err
	}
	return nil
}

func CloseRedis() {
	redisClient.Close()
}
