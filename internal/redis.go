package internal

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Connect to redis server
func ConnectToRedis(ctx *AppContext) error {
	config, err := ctx.GetConfig()
	if err != nil {
		return err
	}

	redisHost, err := config.Get("REDIS_HOST")
	if err != nil {
		return err
	}

	redisPassword, err := config.Get("REDIS_PASSWORD")
	if err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	ctx.SetRedisClient(client)

	return nil
}

// Prefixes the key with the app redis key for namespacing
func RedisKey(ctx *AppContext, key string) (string, error) {
	config, err := ctx.GetConfig()
	if err != nil {
		return "", err
	}

	redisPrefix, err := config.Get("REDIS_PREFIX")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", redisPrefix, key), nil
}
