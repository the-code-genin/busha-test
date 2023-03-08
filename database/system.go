package database

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/the-code-genin/busha-test/internal"
)

const (
	SeededKey = "seeded"
)

type SystemRepository struct {
	ctx         *internal.AppContext
	redisClient *redis.Client
}

// Returns true if the system has been seeded
func (r *SystemRepository) GetSeeded() (bool, error) {
	key, err := internal.RedisKey(r.ctx, SeededKey)
	if err != nil {
		return false, err
	}

	result := r.redisClient.Get(context.Background(), key)
	if result == nil {
		return false, nil
	}

	return result.Val() == "true", nil
}

// Updates the seeded value for the system
func (r *SystemRepository) SetSeeded(val bool) error {
	key, err := internal.RedisKey(r.ctx, SeededKey)
	if err != nil {
		return err
	}

	result := r.redisClient.Set(
		context.Background(),
		key,
		"true",
		0,
	)
	if result == nil {
		return errors.New("an error occured while setting seeded value")
	}

	return nil
}

func NewSystemRepository(ctx *internal.AppContext) (*SystemRepository, error) {
	redisClient, err := ctx.GetRedisClient()
	if err != nil {
		return nil, err
	}

	return &SystemRepository{ctx, redisClient}, nil
}
