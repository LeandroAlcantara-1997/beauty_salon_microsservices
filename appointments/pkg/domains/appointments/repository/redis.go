package repository

import "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(c *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: c,
	}
}
