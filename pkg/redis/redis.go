package redis

import (
	"context"
	"fmt"

	"github.com/pc-configurator/components/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string `envconfig:"REDIS_ADDR" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
}

type Client struct {
	*redis.Client
}

func New(c Config, ctx context.Context) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis.Ping: %w", err)
	}

	if len(pong) == 0 {
		return nil, fmt.Errorf("redis.Ping: result is empty")
	}

	return &Client{Client: client}, nil
}

func (c *Client) Close() {
	err := c.Client.Close()
	if err != nil {
		logger.Error(err, "redis.Close")
	}

	logger.Info("redis: closed")
}
