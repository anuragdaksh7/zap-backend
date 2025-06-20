package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	redisAddr := config.RedisAddr
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: config.RedisPasswd,
		DB:       0,
		PoolSize: 10,
	})

	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Cannot connect to Redis: %v", err)
	}
	log.Println("Redis Connected!")
}
