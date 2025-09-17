package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisClient represents our Redis client.
var RedisClient *redis.Client

// InitRedisClient initializes and returns a new Redis client.
func InitRedisClient(addr string, password string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // Use default DB
	})

	ctx := context.Background()
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	fmt.Println("Successfully connected to Redis!")
	return nil
}
