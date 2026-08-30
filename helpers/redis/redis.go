package redis

import "github.com/redis/go-redis/v9"

type RedisService struct {
}

func New() *RedisService {
	return &RedisService{}
}

func (s *RedisService) Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
