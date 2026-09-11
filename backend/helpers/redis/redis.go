package redis

import (
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

type RedisService struct {
}

func New() *RedisService {
	return &RedisService{}
}

func (s *RedisService) Client() *redis.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	url := os.Getenv("REDIS_URL")
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	return rdb
}

func (s *RedisService) Close(client *redis.Client) {
	if err := client.Close(); err != nil {
		log.Panicln("close redis: %w", err)
	}
	log.Println("Redis connection closed")
}
