package internal

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Connect to redis server
func ConnectToRedis() (*redis.Client, error) {
	redisHost, err := DefaultConfig.Get("REDIS_HOST")
	if err != nil {
		return nil, err
	}

	redisPassword, err := DefaultConfig.Get("REDIS_PASSWORD")
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	return client, nil
}

// Prefixes the key with the app redis key for namespacing
func RedisKey(key string) string {
	redisPrefix, err := DefaultConfig.Get("REDIS_PREFIX")
	if err != nil {
		return key
	}
	return fmt.Sprintf("%s.%s", redisPrefix, key)
}
