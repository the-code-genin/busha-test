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

	redisHost, err := config.GetRedisHost()
	if err != nil {
		return err
	}

	redisPassword, err := config.GetRedisPassword()
	if err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	ctx.setRedisClient(client)

	return nil
}

// Prefixes the key with the app redis key for namespacing
func RedisKey(ctx *AppContext, key string) (string, error) {
	config, err := ctx.GetConfig()
	if err != nil {
		return "", err
	}

	redisPrefix, err := config.GetRedisPrefix()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", redisPrefix, key), nil
}
