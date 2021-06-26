package redis

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/go-redis/redis/v8"
	"time"
)

// Returns a new redis client
func NewRedisClient(cfg *config.Config) *redis.Client {
	redisHost := cfg.Redis.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisHost,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize: cfg.Redis.PoolSize,
		PoolTimeout: time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password: cfg.Redis.Password,
		DB: cfg.Redis.DB,
	})

	return client
}