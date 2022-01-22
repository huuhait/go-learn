package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client  *redis.Client
	context context.Context
}

func NewRedisClient(address string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	return &RedisClient{
		client:  client,
		context: context.Background(),
	}
}

func (c *RedisClient) Set(key string, value interface{}) error {
	_, err := c.client.Set(c.context, key, value, 0).Result()

	return err
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.client.Get(c.context, key).Result()
}

func (c *RedisClient) Delete(key string) error {
	_, err := c.client.Del(c.context, key).Result()

	return err
}
