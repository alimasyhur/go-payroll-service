package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, val []byte, expiration time.Duration) (err error)
	Get(ctx context.Context, key string) (result string, err error)
	Del(ctx context.Context, key string) (err error)
}

type client struct {
	redisKeyPrefix string
	redisClient    *redis.Client
}

func New(cfg Config) RedisClient {
	fmt.Println("Try NewRedis ...")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       0,
	})

	fmt.Println("Redis", redisClient.Ping(context.Background()))

	return &client{
		redisKeyPrefix: cfg.Prefix,
		redisClient:    redisClient,
	}
}

func (c *client) Set(ctx context.Context, key string, val []byte, expiration time.Duration) (err error) {
	key = fmt.Sprintf("%s-%s", c.redisKeyPrefix, key)
	err = c.redisClient.Set(ctx, key, string(val), expiration).Err()
	return
}

func (c *client) Get(ctx context.Context, key string) (result string, err error) {
	key = fmt.Sprintf("%s-%s", c.redisKeyPrefix, key)
	result, err = c.redisClient.Get(ctx, key).Result()
	return
}

func (c *client) Del(ctx context.Context, key string) (err error) {
	key = fmt.Sprintf("%s-%s", c.redisKeyPrefix, key)
	err = c.redisClient.Del(ctx, key).Err()
	return
}
