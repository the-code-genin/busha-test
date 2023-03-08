package repositories

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
	redisClient *redis.Client
}

// Returns true if the system has been seeded
func (r *SystemRepository) GetSeeded() (bool, error) {
	result := r.redisClient.Get(context.Background(), internal.RedisKey(SeededKey))
	if result == nil {
		return false, nil
	}
	return result.Val() == "true", nil
}

// Updates the seeded value for the system
func (r *SystemRepository) SetSeeded(val bool) error {
	result := r.redisClient.SetNX(
		context.Background(),
		internal.RedisKey(SeededKey),
		"true",
		0,
	)
	if result == nil {
		return errors.New("an error occured while setting seeded value")
	}
	return nil
}

func NewSystemRepository(ctx context.Context) (*SystemRepository, error) {
	redisClient, err := internal.GetRedisClient(ctx)
	if err != nil {
		return nil, err
	}

	return &SystemRepository{redisClient}, nil
}
