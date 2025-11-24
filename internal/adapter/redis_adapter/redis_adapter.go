package redis_adapter

import "github.com/pc-configurator/components/pkg/redis"

type RedisAdapter struct {
	client *redis.Client
}

func New(c *redis.Client) *RedisAdapter {
	return &RedisAdapter{client: c}
}
