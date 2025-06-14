package config

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Check connection
	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}
