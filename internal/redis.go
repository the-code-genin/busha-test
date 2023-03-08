package internal

import (
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
