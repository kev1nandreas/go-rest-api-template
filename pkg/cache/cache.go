package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kev1nandreas/go-rest-api-template/env"
)

type Cache interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Keys(context.Context, string) *redis.StringSliceCmd
	Del(context.Context, ...string) *redis.IntCmd
}

func NewRedisClient() *redis.Client {
	redis_host := env.GetEnvString("REDIS_HOST", "localhost")
	redis_port := env.GetEnvString("REDIS_PORT", "6379")
	redis_password := env.GetEnvString("REDIS_PASSWORD", "")

	return redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port, // Redis server address (change to localhost when running local)
		Password: redis_password,                // Password, leave empty if none
		DB:       0,                             // Default DB
	})
}
